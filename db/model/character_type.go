package model

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type CharacterType string

const (
	CharacterMain CharacterType = "Main"
	CharacterNPC                = "NPC"
)

var characterTypeVariants = []CharacterType{
	CharacterMain,
	CharacterNPC,
}

var _ sql.Scanner = (*CharacterType)(nil)
var _ driver.Valuer = (*CharacterType)(nil)

func (m *CharacterType) Scan(value any) (err error) {
	if *m, err = scanEnum(value, characterTypeVariants); err != nil {
		return fmt.Errorf("failed to scan CharacterType: %w", err)
	}
	return nil
}

func (m CharacterType) Value() (driver.Value, error) {
	value, err := valueEnum(m, characterTypeVariants)
	if err != nil {
		return nil, fmt.Errorf("failed to value CharacterType: %w", err)
	}
	return value, nil
}
