package saves

import (
	"fmt"
	"slices"

	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
)

type Model struct {
	*backbone.BaseModel
	List []*model.GameSave
}

func NewModel() *Model {
	return &Model{
		BaseModel: backbone.NewBaseModel(),
	}
}

func (m *Model) Load() error {
	err := db.DB().
		NewSelect().
		Model(&m.List).
		Order("created_at desc").
		Scan(m.Ctx)

	if err != nil {
		return fmt.Errorf("failed to load saves: %w", err)
	}

	return nil
}

func (m *Model) Delete(deletingSave *model.GameSave) error {
	_, err := db.DB().
		NewDelete().
		Model(deletingSave).
		WherePK().
		Exec(m.Ctx)

	if err != nil {
		return fmt.Errorf("failed to delete game save: %w", err)
	}

	m.List = slices.DeleteFunc(m.List, func(save *model.GameSave) bool {
		return deletingSave.ID == save.ID
	})

	return nil
}
