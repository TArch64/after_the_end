package appwindow

import (
	"after_the_end/app/game"
	"after_the_end/app/gamewizard"
	"after_the_end/app/router"
	"after_the_end/app/saves"
	"after_the_end/app/start"
	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatelessView
	window *qt.QMainWindow
}

func NewWindowView() *View {
	return &View{
		StatelessView: backbone.NewStatelessView(),
	}
}

func (v *View) ViewInit(_ *qt.QWidget) {
	v.window = qt.NewQMainWindow2()
	v.window.SetObjectName("main_window")
	v.window.SetWindowTitle("AfterTheEnd")
	v.window.ResizeWithQSize(qt.QGuiApplication_PrimaryScreen().Geometry().Size())

	centralWidget := v.Mount(router.NewView(&router.Options{
		InitialRoute: router.RouteStart,

		Routes: router.Routes{
			router.RouteStart:      start.NewView(),
			router.RouteSaves:      saves.NewView(),
			router.RouteGameWizard: gamewizard.NewView(),
			router.RouteGame:       game.NewView(),
		},
	}))

	v.window.SetCentralWidget(centralWidget)
	v.window.ShowMaximized()
}
