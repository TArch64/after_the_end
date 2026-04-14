package gamewizard

import (
	"database/sql"
	"errors"
	"fmt"

	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"

	"github.com/uptrace/bun"
)

type Model struct {
	*backbone.BaseModel
	GameSave      *model.GameSave
	MainCharacter *model.Character
}

func NewModel() *Model {
	return &Model{
		BaseModel: backbone.NewBaseModel(),
	}
}

func (m *Model) Load(gameSave *model.GameSave) error {
	m.GameSave = gameSave
	return m.loadMainCharacter()
}

func (m *Model) loadMainCharacter() error {
	m.MainCharacter = &model.Character{}

	err := db.DB().
		NewSelect().
		Model(m.MainCharacter).
		Where("save_id = ?", m.GameSave.ID).
		Where("type = ?", model.CharacterMain).
		Limit(1).
		Scan(m.Ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return m.createMainCharacter()
	}

	if err != nil {
		return fmt.Errorf("load main character: %w", err)
	}

	return nil
}

func (m *Model) createMainCharacter() error {
	m.MainCharacter.SaveID = m.GameSave.ID
	m.MainCharacter.Type = model.CharacterMain

	_, err := db.DB().
		NewInsert().
		Model(m.MainCharacter).
		Exec(m.Ctx)

	if err != nil {
		return fmt.Errorf("create main character: %w", err)
	}

	return nil
}

func (m *Model) UpdateMainCharacter(patch map[string]any) error {
	query := db.DB().
		NewUpdate().
		Model((*model.Character)(nil)).
		Where("id = ?", m.MainCharacter.ID)

	for field, value := range patch {
		query = query.Set("? = ?", bun.Ident(field), value)
	}

	if _, err := query.Exec(m.Ctx); err != nil {
		return fmt.Errorf("update main character: %w", err)
	}

	return nil
}
