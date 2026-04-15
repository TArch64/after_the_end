package maincolumn

import (
	"after_the_end/helper/qtgeometry"

	"github.com/mappu/miqt/qt"
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
	w.Container = qt.NewQWidget(w.target)
	w.Container.SetObjectName("main_column_container")
	w.Container.SetLayout(w.QLayout)
	qtgeometry.Read(w.target, w.renderSize)

}

func (w *Widget) renderSize(geometry *qt.QRect) {
	width := min(1000, int(float32(geometry.Width())*0.6))
	x := max(0, (geometry.Width()-width)/2)
	w.Container.SetGeometry(x, 0, width, geometry.Height())
}
