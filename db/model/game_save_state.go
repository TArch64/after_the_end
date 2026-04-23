package model

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type GameSaveState string

const (
	GameSaveNew           GameSaveState = "New"
	GameSaveGeneratingMap GameSaveState = "GeneratingMap"
	GameSaveReady         GameSaveState = "Ready"
)

var GameSaveStateVariants = []GameSaveState{
	GameSaveNew,
	GameSaveGeneratingMap,
	GameSaveReady,
}

var _ sql.Scanner = (*GameSaveState)(nil)
var _ driver.Valuer = (*GameSaveState)(nil)

func (m *GameSaveState) Scan(value any) (err error) {
	if *m, err = scanEnum(value, GameSaveStateVariants); err != nil {
		return fmt.Errorf("failed to scan GameSaveState: %w", err)
	}
	return nil
}

func (m GameSaveState) Value() (driver.Value, error) {
	value, err := valueEnum(m, GameSaveStateVariants)
	if err != nil {
		return nil, fmt.Errorf("failed to value GameSaveState: %w", err)
	}
	return value, nil
}
