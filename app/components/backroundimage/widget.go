package backroundimage

import (
	"fmt"

	qt "github.com/mappu/miqt/qt6"
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
	w.SetProperty("background-image", qt.NewQVariant8(true))
	w.SetStyleSheet(fmt.Sprintf("[background-image='true'] { background: url(%s) }", w.src))

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
	w.overlay.SetProperty("background-image-overlay", qt.NewQVariant8(true))
	w.overlay.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Expanding)
	w.overlay.SetStyleSheet(fmt.Sprintf("[background-image-overlay='true'] { background: %s }", w.overlayColor))
	return w.overlay
}

func (w *Widget) renderContent() *qt.QWidget {
	w.Content = qt.NewQWidget2()
	w.Content.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Expanding)
	return w.Content
}
