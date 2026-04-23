package backroundimage

import (
	"fmt"

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
	Src          string
	OverlayColor string
}

func New(options *Options) *Widget {
	widget := &Widget{
		QWidget:      qt.NewQWidget2(),
		src:          options.Src,
		overlayColor: options.OverlayColor,
	}

	widget.render()
	return widget
}

func (w *Widget) render() {
	w.SetProperty("background-image", qt.NewQVariant11(true))
	w.SetStyleSheet(fmt.Sprintf("QWidget[background-image='true'] { background: url(%s) }", w.src))

	layout := qt.NewQVBoxLayout(w.QWidget)
	layout.SetContentsMargins(0, 0, 0, 0)

	if w.overlayColor != "" {
		layout.AddWidget(w.renderOverlay())
		layout = qt.NewQVBoxLayout(w.overlay)
		layout.SetContentsMargins(0, 0, 0, 0)
	}

	layout.AddWidget(w.renderContent())
}

func (w *Widget) renderOverlay() *qt.QWidget {
	w.overlay = qt.NewQWidget(w.QWidget)
	w.overlay.SetObjectName("background_overlay")
	w.overlay.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Expanding)
	w.overlay.SetStyleSheet(fmt.Sprintf("#background_overlay { background: %s }", w.overlayColor))
	return w.overlay
}

func (w *Widget) renderContent() *qt.QWidget {
	w.Content = qt.NewQWidget2()
	w.Content.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Expanding)
	w.Content.SetObjectName("background_content")
	return w.Content
}
