package gamewizard

import (
	"after_the_end/backbone"
	"after_the_end/db/model"
)

type Model struct {
	*backbone.BaseModel
	GameSave *model.GameSave
}

func NewModel() *Model {
	return &Model{
		BaseModel: backbone.NewBaseModel(),
	}
}

func (v *Model) Load(gameSave *model.GameSave) {
	v.GameSave = gameSave
}
