package saves

import (
	"log/slog"
	"slices"

	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
	"after_the_end/logs"
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

func (m *Model) Load() {
	err := db.DB().
		NewSelect().
		Model(&m.List).
		Order("created_at desc").
		Scan(m.Ctx)

	if err != nil {
		slog.Error("failed to load saves",
			logs.AttrError(err),
		)
	}
}

func (m *Model) Delete(deletingSave *model.GameSave) bool {
	_, err := db.DB().
		NewDelete().
		Model(deletingSave).
		WherePK().
		Exec(m.Ctx)

	if err != nil {
		slog.Error("failed to delete save",
			logs.AttrError(err),
		)

		return false
	}

	m.List = slices.DeleteFunc(m.List, func(save *model.GameSave) bool {
		return deletingSave.ID == save.ID
	})

	return true
}
