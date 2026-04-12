package saves

import (
	"fmt"

	"after_the_end/backbone"
	"after_the_end/db/model"
	"after_the_end/helper/dateformat"
)

type SaveModel struct {
	*backbone.BaseModel
	parent   *Model
	GameSave *model.GameSave
}

func NewSaveModel(parent *Model, gameSave *model.GameSave) *SaveModel {
	return &SaveModel{
		BaseModel: backbone.NewBaseModel(),
		parent:    parent,
		GameSave:  gameSave,
	}
}

func (m *SaveModel) FormatTitle() string {
	return fmt.Sprintf("Save #%d", m.GameSave.ID)
}

func (m *SaveModel) FormatUpdatedAt() string {
	return dateformat.Relative(m.GameSave.UpdatedAt)
}
