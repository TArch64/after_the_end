package maincolumn

import (
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
	w.renderSize()

	w.target.OnResizeEvent(func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent) {
		super(event)
		w.renderSize()
	})
}

func (w *Widget) renderSize() {
	parent := w.target.Geometry()
	width := min(1000, int(float32(parent.Width())*0.6))
	left := max(0, (parent.Width()-width)/2)
	geometry := qt.NewQRect4(left, 0, width, parent.Height())
	w.Container.SetGeometryWithGeometry(geometry)
}
