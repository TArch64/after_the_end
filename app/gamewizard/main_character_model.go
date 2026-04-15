package gamewizard

import (
	"database/sql"
	"errors"
	"fmt"

	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
)

type MainCharacterModel struct {
	*backbone.BaseModel
	Character *model.Character
}

func NewMainCharacterModel() *MainCharacterModel {
	return &MainCharacterModel{
		BaseModel: backbone.NewBaseModel(),
	}
}

func (m *MainCharacterModel) Load(gameSaveID model.ID) error {
	m.Character = &model.Character{}

	err := db.DB().
		NewSelect().
		Model(m.Character).
		Where("save_id = ?", gameSaveID).
		Where("type = ?", model.CharacterMain).
		Limit(1).
		Scan(m.Ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return m.create(gameSaveID)
	}

	if err != nil {
		return fmt.Errorf("load main character: %w", err)
	}

	return nil
}

func (m *MainCharacterModel) create(gameSaveID model.ID) error {
	m.Character.SaveID = gameSaveID
	m.Character.Type = model.CharacterMain

	_, err := db.DB().
		NewInsert().
		Model(m.Character).
		Exec(m.Ctx)

	if err != nil {
		return fmt.Errorf("create main character: %w", err)
	}

	return nil
}

func (m *MainCharacterModel) Save(columns ...string) error {
	_, err := db.DB().
		NewUpdate().
		Model(m.Character).
		Column(columns...).
		WherePK().
		Exec(m.Ctx)

	if err != nil {
		return fmt.Errorf("save main character: %w", err)
	}

	return nil
}
