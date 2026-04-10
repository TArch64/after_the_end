package start

import (
	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.BaseView
}

func NewView() *View {
	return &View{
		BaseView: backbone.NewBaseView(),
	}
}

func (v *View) ViewInit(parent *qt.QWidget) {
	widget := qt.NewQWidget2()
	widget.SetStyleSheet("background-image: url(:/images/background.jpg)")

	row := qt.NewQHBoxLayout2()
	row.AddWidget3(v.renderAside(), 0, qt.AlignVCenter)
	row.AddStretch()

	widget.SetLayout(row.QLayout)

	cover := qt.NewQVBoxLayout(parent)
	cover.SetContentsMargins(0, 0, 0, 0)
	cover.AddWidget(widget)
}

func (v *View) renderAside() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetStyleSheet("background: transparent")

	layout := qt.NewQVBoxLayout2()
	layout.SetContentsMargins(100, 20, 100, 20)
	layout.AddStretch()
	layout.AddWidget(v.renderTitle())
	layout.AddStretch()

	widget.SetLayout(layout.QLayout)
	return widget
}

func (v *View) renderTitle() *qt.QWidget {
	title := qt.NewQLabel3("After the End")
	title.SetAlignment(qt.AlignCenter)

	title.SetStyleSheet(`
		color: #fff;
		font-size: 80px;
		font-weight: 400;
		font-family: 'Pixelify Sans';
		background: transparent;`)

	glow := qt.NewQGraphicsDropShadowEffect2(title.QObject)
	glow.SetBlurRadius(20)
	glow.SetColor(qt.NewQColor11(0, 0, 0, 120))
	glow.SetOffset2(2, 4)
	title.SetGraphicsEffect(glow.QGraphicsEffect)

	return title.QWidget
}
