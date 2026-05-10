package model

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"

	"after_the_end/helper/axial"
)

type AxialCoord axial.Coord

var _ sql.Scanner = (*AxialCoord)(nil)
var _ driver.Valuer = (*AxialCoord)(nil)

func (c *AxialCoord) Scan(value any) error {
	if value == nil {
		*c = AxialCoord{}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan AxialCoord: value is not a string")
	}

	coord, err := axial.ParseCoord(str)
	if err != nil {
		return fmt.Errorf("failed to scan AxialCoord: %w", err)
	}

	*c = AxialCoord(*coord)
	return nil
}

func (c AxialCoord) Value() (driver.Value, error) {
	return axial.Coord(c).StringKey(), nil
}

func (c *AxialCoord) S() int {
	return axial.Coord(*c).S()
}

func (c *AxialCoord) StringKey() string {
	return axial.Coord(*c).StringKey()
}
