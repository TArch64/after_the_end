package axial

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
)

var _ sql.Scanner = (*Coord)(nil)
var _ driver.Valuer = (*Coord)(nil)

func (c *Coord) Scan(value any) error {
	if value == nil {
		*c = Coord{}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan Coord: value is not a string")
	}

	coord, err := ParseCoord(str)
	if err != nil {
		return fmt.Errorf("failed to scan Coord: %w", err)
	}

	*c = *coord
	return nil
}

func (c *Coord) Value() (driver.Value, error) {
	return c.StringKey(), nil
}
