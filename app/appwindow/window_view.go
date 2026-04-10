package appwindow

import (
	"after_the_end/app/appwindow/start"
	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type WindowView struct {
	*backbone.BaseView
	window *qt.QMainWindow
}

func NewWindowView() *WindowView {
	return &WindowView{
		BaseView: backbone.NewBaseView(),
	}
}

func (v *WindowView) ViewInit(_ *qt.QWidget) {
	v.window = qt.NewQMainWindow2()
	v.window.SetWindowTitle("AfterTheEnd")

	centralWidget := qt.NewQWidget2()
	v.Mount(centralWidget, start.NewView())

	v.window.SetCentralWidget(centralWidget)
	v.window.ShowMaximized()
}
