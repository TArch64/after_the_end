package start

import (
	"log/slog"

	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
	"after_the_end/logs"
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

func (m *MenuModel) CountSaves() {
	savesCount, err := db.DB().
		NewSelect().
		Model((*model.GameSave)(nil)).
		Count(m.Ctx)

	if err != nil {
		slog.Error("failed to count game saves",
			logs.AttrError(err),
		)
	}

	m.SavesCount = uint8(savesCount)
}

func (m *MenuModel) NewGame() {
	save := &model.GameSave{
		Position: m.SavesCount,
	}

	_, err := db.DB().
		NewInsert().
		Model(save).
		Exec(m.Ctx)

	if err == nil {
		m.SavesCount++
	} else {
		slog.Error("failed to create new game save",
			logs.AttrError(err),
		)
	}
}
