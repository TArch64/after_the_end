package start

import (
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatelessView
	layout *qt.QLayout
}

func NewView() *View {
	return &View{
		StatelessView: backbone.NewStatelessView(),
	}
}

func (v *View) Layout() *qt.QLayout {
	return v.layout
}

func (v *View) ViewInit(parent *qt.QWidget) {
	widget := qt.NewQWidget2()
	widget.SetObjectName("start_window")
	widget.SetStyleSheet("#start_window { background: url(:/images/background.jpg) }")

	row := qt.NewQHBoxLayout2()
	row.SetObjectName("start_window_row")
	row.AddWidget3(v.renderAside(), 0, qt.AlignVCenter)
	row.AddStretch()

	widget.SetLayout(row.QLayout)

	cover := qt.NewQVBoxLayout(parent)
	cover.SetObjectName("start_window_cover")
	cover.SetContentsMargins(0, 0, 0, 0)
	cover.AddWidget(widget)
	v.layout = cover.QLayout
}

func (v *View) renderAside() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("aside")

	layout := qt.NewQVBoxLayout2()
	layout.SetObjectName("aside")
	layout.SetContentsMargins(100, 20, 100, 20)

	layout.AddStretch()
	layout.AddWidget(v.renderTitle())
	layout.AddWidget(v.MountForLayout(NewMenuView()))
	layout.AddStretch()

	widget.SetLayout(layout.QLayout)
	return widget
}

func (v *View) renderTitle() *qt.QWidget {
	title := qt.NewQLabel3("After the End")
	title.SetAlignment(qt.AlignCenter)
	title.SetStyleSheet(styled.Title1)
	title.SetGraphicsEffect(styled.TitleShadow())

	return title.QWidget
}
