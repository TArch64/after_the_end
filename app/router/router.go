package router

import (
	"log/slog"

	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.BaseView
	routes        Routes
	currentRoute  backbone.View
	currentLayout *qt.QLayout
	initialRoute  RouteName
	parent        *qt.QWidget
}

type Options struct {
	InitialRoute RouteName
	Routes       Routes
}

func NewView(options *Options) *View {
	return &View{
		BaseView:     backbone.NewBaseView(),
		routes:       options.Routes,
		initialRoute: options.InitialRoute,
	}
}

func (v *View) ViewInit(parent *qt.QWidget) {
	v.parent = parent
	v.renderRoute(v.initialRoute)
	onPush = v.renderRoute
}

func (v *View) renderRoute(name RouteName) {
	newRoute, ok := v.routes[name]
	if !ok {
		slog.Error("unknown route",
			slog.String("name", string(name)),
		)
		return
	}

	if v.currentRoute != nil {
		v.currentRoute.ViewDestroy()
		v.currentLayout.Delete()
	}

	v.currentRoute = newRoute

	widget := qt.NewQWidget2()
	widget.SetObjectName("router_container")
	v.currentRoute.ViewBeforeInit()
	v.currentRoute.ViewInit(widget)

	cover := qt.NewQVBoxLayout(v.parent)
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
