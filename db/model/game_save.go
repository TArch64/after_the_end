package model

import (
	"time"

	"github.com/uptrace/bun"
)

type GameSave struct {
	bun.BaseModel `bun:"table:game_saves,alias:gs"`
	ID            ID            `bun:",pk,autoincrement"`
	Position      uint8         `bun:",notnull"`
	State         GameSaveState `bun:",notnull"`
	CreatedAt     time.Time     `bun:",notnull,default:current_timestamp"`
	UpdatedAt     time.Time     `bun:",notnull,default:current_timestamp"`

	// Relations
	Characters []*Character `bun:"rel:has-many,join:id=save_id"`
}
