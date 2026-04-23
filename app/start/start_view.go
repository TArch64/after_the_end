package start

import (
	"after_the_end/app/components/backroundimage"
	"after_the_end/app/resources"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatefullView[*Model]
}

func NewView() *View {
	return &View{
		StatefullView: backbone.NewStatefullView(NewModel()),
	}
}

func (v *View) ViewBeforeOpen(_ router.Params) error {
	return v.Model.Load()
}

func (v *View) ViewInit() *qt.QWidget {
	widget := backroundimage.New(&backroundimage.Options{
		Src: resources.Image("background.jpg"),
	})

	widget.SetObjectName("start_window")

	row := qt.NewQHBoxLayout(widget.Content)
	row.SetObjectName("start_window_row")
	row.AddWidget3(v.renderAside(), 0, qt.AlignVCenter)
	row.AddStretch()

	return widget.QWidget
}

func (v *View) renderAside() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("aside")

	layout := qt.NewQVBoxLayout2()
	layout.SetObjectName("aside")
	layout.SetContentsMargins(100, 20, 100, 20)

	layout.AddStretch()
	layout.AddWidget(v.renderTitle())
	layout.AddWidget(v.Mount(NewMenuView(v.Model)))
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
