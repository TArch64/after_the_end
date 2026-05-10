package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Character struct {
	bun.BaseModel `bun:"table:characters,alias:c"`
	ID            ID            `bun:",pk,autoincrement"`
	Type          CharacterType `bun:",notnull"`
	Name          string        `bun:",notnull"`
	LocationID    ID            `bun:",notnull"`
	LocationCoord *AxialCoord   `bun:",notnull"`
	SaveID        ID            `bun:",notnull"`
	CreatedAt     time.Time     `bun:",notnull,default:current_timestamp"`
	UpdatedAt     time.Time     `bun:",notnull,default:current_timestamp"`

	// Relations
	Save        *GameSave    `bun:"rel:belongs-to,join:save_id=id"`
	Location    *Location    `bun:"rel:has-one,join:location_id=id"`
	LocationHex *LocationHex `bun:"rel:has-one,join:location_coord=coord"`
}
