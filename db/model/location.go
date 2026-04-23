package model

import (
	"github.com/uptrace/bun"
)

type Location struct {
	bun.BaseModel `bun:"table:locations,alias:l"`
	ID            ID     `bun:",pk,autoincrement"`
	Name          string `bun:",notnull"`
	SaveID        ID     `bun:",notnull"`

	// Relations
	Save  *GameSave      `bun:"rel:belongs-to,join:save_id=id"`
	Hexes []*LocationHex `bun:"rel:has-many,join:id=location_id"`
}
