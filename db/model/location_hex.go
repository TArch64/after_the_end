package model

import (
	"after_the_end/helper/axial"

	"github.com/uptrace/bun"
)

type LocationHex struct {
	bun.BaseModel `bun:"table:location_hexes,alias:lh"`
	*axial.Coord
	Elevation  int `bun:",notnull"`
	LocationID ID  `bun:",notnull"`

	// Relations
	Location *Location `bun:"rel:belongs-to,join:location_id=id"`
}
