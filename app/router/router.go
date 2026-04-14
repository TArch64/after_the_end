package router

import (
	"log/slog"

	"after_the_end/app/dialog/errorreport"
	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatelessView
	routes        Routes
	currentRoute  backbone.View
	currentLayout *qt.QLayout
	initialRoute  Name
	container     *qt.QWidget
}

type Options struct {
	InitialRoute Name
	Routes       Routes
}

func NewView(options *Options) *View {
	return &View{
		StatelessView: backbone.NewStatelessView(),
		routes:        options.Routes,
		initialRoute:  options.InitialRoute,
	}
}

func (v *View) ViewInit() *qt.QWidget {
	v.container = qt.NewQWidget2()
	v.container.SetObjectName("router_container")
	v.renderRoute(v.initialRoute, nil)
	onPush = v.renderRoute
	return v.container
}

func (v *View) renderRoute(name Name, params Params) {
	newRoute, ok := v.routes[name]
	if !ok {
		slog.Error("unknown route",
			slog.String("name", string(name)),
		)
		return
	}

	newRoute.ViewBeforeInit()

	if err := newRoute.ViewBeforeOpen(params); err != nil {
		errorreport.Show(v.Root, err)
		return
	}

	if v.currentRoute != nil {
		v.currentRoute.ViewDestroy()
		v.currentLayout.Delete()
	}

	v.currentRoute = newRoute
	widget := v.Mount(v.currentRoute, true)

	cover := qt.NewQVBoxLayout(v.container)
	cover.SetObjectName("router_container")
	cover.SetContentsMargins(0, 0, 0, 0)
	cover.AddWidget(widget)
	v.currentLayout = cover.QLayout
}

func (v *View) ViewDestroy() {
	if v.currentRoute != nil {
		v.currentRoute.ViewDestroy()
		v.currentLayout.DeleteLater()
	}
}
