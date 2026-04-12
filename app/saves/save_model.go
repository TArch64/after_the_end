package saves

import (
	"after_the_end/backbone"
	"after_the_end/db/model"
	"after_the_end/helper/dateformat"
)

type SaveModel struct {
	*backbone.BaseModel
	GameSave *model.GameSave
}

func NewSaveModel(gameSave *model.GameSave) *SaveModel {
	return &SaveModel{
		BaseModel: backbone.NewBaseModel(),
		GameSave:  gameSave,
	}
}

func (m *SaveModel) FormatUpdatedAt() string {
	return dateformat.Relative(m.GameSave.UpdatedAt)
}
