package gamewizard

import (
	"fmt"

	"after_the_end/app/gamewizard/mapgenerator"
	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
)

type GameSaveModel struct {
	*backbone.BaseModel
	GameSave           *model.GameSave
	MainCharacterModel *MainCharacterModel
}

func NewGameSaveModel() *GameSaveModel {
	return &GameSaveModel{
		BaseModel:          backbone.NewBaseModel(),
		MainCharacterModel: NewMainCharacterModel(),
	}
}

func (m *GameSaveModel) ModelInit() {
	m.BaseModel.ModelInit()
	m.ModelInitChild(m.MainCharacterModel.BaseModel)
}

func (m *GameSaveModel) Load(gameSave *model.GameSave) error {
	m.GameSave = gameSave
	return m.MainCharacterModel.Load(gameSave.ID)
}

func (m *GameSaveModel) Save(columns ...string) error {
	_, err := db.DB().
		NewUpdate().
		Model(m.GameSave).
		Column(columns...).
		WherePK().
		Exec(m.Ctx)

	if err != nil {
		return fmt.Errorf("save game: %w", err)
	}

	return nil
}

func (m *GameSaveModel) GenerateMap() error {
	if err := mapgenerator.Generate(m.Ctx, m.GameSave); err != nil {
		return fmt.Errorf("generate map: %w", err)
	}

	return nil
}
