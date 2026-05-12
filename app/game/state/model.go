package state

import (
	"context"
	"fmt"

	"after_the_end/app/game/command"
	"after_the_end/app/game/command/cmd"
	"after_the_end/app/game/state/pathfinder"
	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
	"after_the_end/helper/axial"
)

type Model struct {
	*backbone.BaseModel
	GameSave       *model.GameSave
	ActiveLocation *model.Location
	MainCharacter  *model.Character
	IsWalking      bool
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

	m.AutoDispose(command.Handle[*cmd.WalkPath](m.walkPath))
	return nil
}

func (m *Model) walkPath(walkCmd *cmd.WalkPath) {
	if m.IsWalking {
		return
	}

	m.IsWalking = true
	var coord *axial.Coord

	finder := pathfinder.New(
		m.ActiveLocation,
		m.MainCharacter.LocationCoord,
		walkCmd.To,
	)

	for _, coord = range finder.Find()[1:] {
		command.BlockingDispatch(cmd.NewMoveMainCharacter(coord))
	}

	m.IsWalking = false
	m.MainCharacter.LocationCoord = coord
	err := m.saveCharacter(m.Ctx, m.MainCharacter, "location_coord")
	if err != nil {
		command.Dispatch(cmd.NewReportErr(err))
		return
	}
}

func (m *Model) saveCharacter(
	ctx context.Context,
	character *model.Character,
	columns ...string,
) error {
	_, err := db.DB().
		NewUpdate().
		Model(character).
		Column(columns...).
		WherePK().
		Exec(ctx)

	return err
}
