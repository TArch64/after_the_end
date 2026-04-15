package backroundimage

import (
	"fmt"

	"after_the_end/helper/qtgeometry"

	"github.com/mappu/miqt/qt"
)

const (
	OverlayDark = "rgba(0, 0, 0, 0.6)"
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

	qtgeometry.Read(w.QWidget, w.resize)
}

func (w *Widget) renderOverlay() {
	w.overlay = qt.NewQWidget(w.QWidget)
	w.overlay.SetObjectName("background_overlay")
	w.overlay.SetStyleSheet(fmt.Sprintf("#background_overlay { background: %s }", w.overlayColor))
}

func (w *Widget) renderContent(container *qt.QWidget) {
	w.Content = qt.NewQWidget(container)
	w.Content.SetObjectName("background_content")
}

func (w *Widget) OnResizeEvent(handler func()) {
	w.onResize = handler
}

func (w *Widget) resize(geometry *qt.QRect) {
	if w.overlay != nil {
		w.overlay.SetGeometryWithGeometry(geometry)
	}

	w.Content.SetGeometryWithGeometry(geometry)

	if w.onResize != nil {
		w.onResize()
	}
}
