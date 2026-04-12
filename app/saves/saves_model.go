package saves

import (
	"log/slog"

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

func (v *Model) Load() {
	err := db.DB().
		NewSelect().
		Model(&v.List).
		Scan(v.Ctx)

	if err != nil {
		slog.Error("failed to load saves",
			logs.AttrError(err),
		)
	}
}
