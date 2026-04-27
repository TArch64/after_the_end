package maincolumn

import (
	qt "github.com/mappu/miqt/qt6"
)

type Widget struct {
	*qt.QVBoxLayout
	target    *qt.QWidget
	Container *qt.QWidget
}

func New(target *qt.QWidget) *Widget {
	widget := &Widget{
		QVBoxLayout: qt.NewQVBoxLayout2(),
		target:      target,
	}
	widget.render()
	return widget
}

func (w *Widget) render() {
	row := qt.NewQHBoxLayout(w.target)
	row.SetContentsMargins(50, 50, 50, 50)
	row.AddWidget(w.renderContainer())
}

func (w *Widget) renderContainer() *qt.QWidget {
	w.Container = qt.NewQWidget2()
	w.Container.SetLayout(w.QLayout)
	w.Container.SetMaximumWidth(1000)
	w.Container.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Expanding)
	return w.Container
}
