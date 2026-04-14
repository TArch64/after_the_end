package gamewizard

import (
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
)

type CreateCharacterView struct {
	*backbone.StatelessView
	model *Model
}

func NewCreateCharacterView(model *Model) *CreateCharacterView {
	return &CreateCharacterView{
		StatelessView: backbone.NewStatelessView(),
		model:         model,
	}
}

func (v *CreateCharacterView) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("create_character")

	column := qt.NewQVBoxLayout(widget)
	column.SetObjectName("create_character")

	column.AddStretch()
	column.AddWidget3(v.renderTitle(), 0, qt.AlignCenter)
	column.AddStretch()

	return widget
}

func (v *CreateCharacterView) renderTitle() *qt.QWidget {
	title := qt.NewQLabel3("Name Your Stranger")
	title.SetObjectName("create_character_title")
	title.SetStyleSheet(styled.S("#create_character_title", styled.Title2))
	title.SetGraphicsEffect(styled.TitleShadow())
	return title.QWidget
}
