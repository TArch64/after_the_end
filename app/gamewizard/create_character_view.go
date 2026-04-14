package gamewizard

import (
	"after_the_end/backbone"

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
	qt.NewQLabel5("create_character_view", widget)
	return widget
}
