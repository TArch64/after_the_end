package start

import (
	"fmt"

	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
)

type Model struct {
	*backbone.BaseModel
	SavesCount uint8
}

func NewModel() *Model {
	return &Model{
		BaseModel: backbone.NewBaseModel(),
	}
}

func (m *Model) Load() error {
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

func (m *Model) NewGame() (*model.GameSave, error) {
	gameSave := &model.GameSave{
		Position: m.SavesCount,
		State:    model.GameSaveCreateMainCharacter,
	}

	_, err := db.DB().
		NewInsert().
		Model(gameSave).
		Exec(m.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to create new game save: %w", err)
	}

	m.SavesCount++
	return gameSave, nil
}

func (m *Model) GetLastGame() (*model.GameSave, error) {
	gameSave := &model.GameSave{}

	err := db.DB().
		NewSelect().
		Model(gameSave).
		Limit(1).
		OrderExpr("updated_at DESC").
		Scan(m.Ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to get last game save: %w", err)
	}

	return gameSave, nil
}
