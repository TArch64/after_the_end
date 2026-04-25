package axial

import (
	"fmt"
)

type Coord struct {
	Q int `json:"q" bun:",pk,notnull"`
	R int `json:"r" bun:",pk,notnull"`
	S int `json:"s" bun:",notnull"`
}

func (c Coord) StringKey() string {
	return fmt.Sprintf("q%dr%ds%d", c.Q, c.R, c.S)
}
