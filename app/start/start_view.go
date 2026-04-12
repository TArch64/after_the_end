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

func (v *View) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("start_window")
	widget.SetStyleSheet("#start_window { background: url(:/images/background.jpg) }")

	row := qt.NewQHBoxLayout2()
	row.SetObjectName("start_window_row")
	row.AddWidget3(v.renderAside(), 0, qt.AlignVCenter)
	row.AddStretch()

	widget.SetLayout(row.QLayout)
	return widget
}

func (v *View) renderAside() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("aside")

	layout := qt.NewQVBoxLayout2()
	layout.SetObjectName("aside")
	layout.SetContentsMargins(100, 20, 100, 20)

	layout.AddStretch()
	layout.AddWidget(v.renderTitle())
	layout.AddWidget(v.Mount(NewMenuView()))
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
