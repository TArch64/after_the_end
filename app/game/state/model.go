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
	MainCharacter  *model.Character
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

	for _, character := range m.GameSave.Characters {
		if character.Type == model.CharacterMain {
			m.MainCharacter = character
			break
		}
	}

	for _, location := range m.GameSave.Locations {
		if m.MainCharacter.LocationID == location.ID {
			m.ActiveLocation = location
			break
		}
	}

	return nil
}
