package model

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

type CharacterType uint8

const (
	CharacterMain CharacterType = iota
	CharacterNPC
)

var _ sql.Scanner = (*CharacterType)(nil)

func (m *CharacterType) Scan(value any) error {
	if num, ok := value.(int64); ok && m.validateInt(num) {
		*m = CharacterType(num)
		return nil
	}
	return errors.New("scan: invalid CharacterType")
}

func (m CharacterType) Value() (driver.Value, error) {
	if num := int64(m); m.validateInt(num) {
		return num, nil
	}
	return nil, errors.New("value: invalid CharacterType")
}

func (m CharacterType) validateInt(num int64) bool {
	value := CharacterType(num)
	return value == CharacterNPC || value == CharacterMain
}
