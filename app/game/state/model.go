package state

import (
	"fmt"

	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
)

type Model struct {
	*backbone.BaseModel
	GameSave       *model.GameSave
	ActiveLocation *model.Location
}

func NewModel() *Model {
	return &Model{
		BaseModel: backbone.NewBaseModel(),
	}
}

func (m *Model) Load(gameSave *model.GameSave) error {
	m.GameSave = gameSave

	err := db.DB().
		NewSelect().
		Model(m.GameSave).
		WherePK().
		Relation("Characters").
		Relation("Locations").
		Relation("Locations.Hexes").
		Scan(m.Ctx)

	if err != nil {
		return fmt.Errorf("load game data: %w", err)
	}

	m.ActiveLocation = m.GameSave.Locations[0]
	return nil
}
