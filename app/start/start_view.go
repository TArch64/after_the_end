package start

import (
	"after_the_end/app/components/backroundimage"
	"after_the_end/app/resources"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	qt "github.com/mappu/miqt/qt6"
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

	row := qt.NewQHBoxLayout(widget.Content)
	row.AddWidget3(v.renderAside(), 0, qt.AlignVCenter)
	row.AddStretch()

	return widget.QWidget
}

func (v *View) renderAside() *qt.QWidget {
	widget := qt.NewQWidget2()

	layout := qt.NewQVBoxLayout2()
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
	title.SetProperty("text-title", qt.NewQVariant4(1))
	title.SetGraphicsEffect(styled.TitleShadow())

	return title.QWidget
}
