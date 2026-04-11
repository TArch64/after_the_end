package saves

import (
	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.BaseView
	layout *qt.QLayout
}

func NewView() *View {
	return &View{
		BaseView: backbone.NewBaseView(),
	}
}

func (v *View) Layout() *qt.QLayout {
	return v.layout
}

func (v *View) ViewInit(parent *qt.QWidget) {
	widget := qt.NewQWidget2()
	widget.SetObjectName("saves")
	widget.SetStyleSheet("background-image: url(:/images/background.jpg)")

	column := qt.NewQVBoxLayout2()
	column.SetObjectName("saves_column")

	column.AddStretch()
	column.AddWidget3(v.renderContainer(), 0, qt.AlignCenter)
	column.AddStretch()

	widget.SetLayout(column.QLayout)

	cover := qt.NewQVBoxLayout(parent)
	cover.SetObjectName("start_window_cover")
	cover.SetContentsMargins(0, 0, 0, 0)
	cover.AddWidget(widget)
	v.layout = cover.QLayout
}

func (v *View) renderContainer() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("saves_container")

	screen := qt.QGuiApplication_PrimaryScreen().Geometry()
	width := max(float32(screen.Width())*0.6, 1000)
	height := max(float32(screen.Height())*0.6, 1000)
	widget.SetFixedSize2(int(width), int(height))

	widget.SetStyleSheet(`
		background-color: #dddddd;
		color: #444444;
		font-size: 18px;
		font-weight: bold;
		padding: 12px 32px;
		border: 3px solid #ffffff;
		border-right-color: #aaaaaa;
		border-bottom-color: #aaaaaa;`)

	return widget
}
