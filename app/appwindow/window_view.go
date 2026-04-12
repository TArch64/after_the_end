package appwindow

import (
	"after_the_end/app/router"
	"after_the_end/app/saves"
	"after_the_end/app/start"
	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type WindowView struct {
	*backbone.StatelessView
	window *qt.QMainWindow
}

func NewWindowView() *WindowView {
	return &WindowView{
		StatelessView: backbone.NewStatelessView(),
	}
}

func (v *WindowView) ViewInit(_ *qt.QWidget) {
	v.window = qt.NewQMainWindow2()
	v.window.SetObjectName("main_window")
	v.window.SetWindowTitle("AfterTheEnd")

	centralWidget := qt.NewQWidget2()
	centralWidget.SetObjectName("window_central")

	v.MountToWidget(centralWidget, router.NewView(&router.Options{
		InitialRoute: router.RouteStart,

		Routes: router.Routes{
			router.RouteStart: start.NewView(),
			router.RouteSaves: saves.NewView(),
		},
	}))

	v.window.SetCentralWidget(centralWidget)
	v.window.ShowMaximized()
}
