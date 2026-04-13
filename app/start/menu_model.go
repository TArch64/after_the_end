package start

import (
	"fmt"

	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
)

type MenuModel struct {
	*backbone.BaseModel
	SavesCount uint8
}

func NewMenuModel() *MenuModel {
	return &MenuModel{
		BaseModel: backbone.NewBaseModel(),
	}
}

func (m *MenuModel) CountSaves() error {
	savesCount, err := db.DB().
		NewSelect().
		Model((*model.GameSave)(nil)).
		Count(m.Ctx)

	if err != nil {
		return fmt.Errorf("failed to count saves: %w", err)
	}

	m.SavesCount = uint8(savesCount)
	return nil
}

func (m *MenuModel) NewGame() error {
	save := &model.GameSave{
		Position: m.SavesCount,
		State:    model.GameSaveCreateMainCharacter,
	}

	_, err := db.DB().
		NewInsert().
		Model(save).
		Exec(m.Ctx)

	if err != nil {
		return fmt.Errorf("failed to create new game save: %w", err)
	}

	m.SavesCount++
	return nil
}
