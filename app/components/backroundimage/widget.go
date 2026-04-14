package backroundimage

import (
	"fmt"

	"github.com/mappu/miqt/qt"
)

type Widget struct {
	*qt.QWidget
	src          string
	overlayColor string
	overlay      *qt.QWidget
	onResize     func()
	Content      *qt.QWidget
}

type Options struct {
	Parent       *qt.QWidget
	Src          string
	OverlayColor string
}

func New(options *Options) *Widget {
	var qWidget *qt.QWidget
	if options.Parent == nil {
		qWidget = qt.NewQWidget2()
	} else {
		qWidget = qt.NewQWidget(options.Parent)
	}

	widget := &Widget{
		QWidget:      qWidget,
		src:          options.Src,
		overlayColor: options.OverlayColor,
	}

	widget.render()
	return widget
}

func (w *Widget) render() {
	w.SetProperty("background-image", qt.NewQVariant11(true))
	w.SetStyleSheet(fmt.Sprintf("QWidget[background-image='true'] { background: url(%s) }", w.src))

	if w.overlayColor == "" {
		w.renderContent(w.QWidget)
	} else {
		w.renderOverlay()
		w.renderContent(w.overlay)
	}

	w.QWidget.OnResizeEvent(w.onResizeEvent)
}

func (w *Widget) renderOverlay() {
	w.overlay = qt.NewQWidget(w.QWidget)
	w.overlay.SetObjectName("background_overlay")
	w.overlay.SetStyleSheet(fmt.Sprintf("#background_overlay { background: %s }", w.overlayColor))
	w.overlay.SetGeometryWithGeometry(w.Geometry())
}

func (w *Widget) renderContent(container *qt.QWidget) {
	w.Content = qt.NewQWidget(container)
	w.Content.SetObjectName("background_content")
	w.Content.SetGeometryWithGeometry(container.Geometry())
}

func (w *Widget) OnResizeEvent(handler func()) {
	w.onResize = handler
}

func (w *Widget) onResizeEvent(super func(event *qt.QResizeEvent), event *qt.QResizeEvent) {
	super(event)
	geometry := w.Geometry()

	if w.overlay != nil {
		w.overlay.SetGeometryWithGeometry(geometry)
	}

	w.Content.SetGeometryWithGeometry(geometry)

	if w.onResize != nil {
		w.onResize()
	}
}
