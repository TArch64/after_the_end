package model

import (
	"github.com/uptrace/bun"
)

type LocationHex struct {
	bun.BaseModel `bun:"table:location_hexes,alias:lh"`
	Coord         *AxialCoord `bun:",notnull"`
	Elevation     int         `bun:",notnull"`
	LocationID    ID          `bun:",notnull"`

	// Relations
	Location *Location `bun:"rel:belongs-to,join:location_id=id"`
}
