package model

import (
	"time"

	"github.com/uptrace/bun"
)

type GameSave struct {
	bun.BaseModel `bun:"table:game_saves,alias:gs"`
	ID            ID        `bun:",pk,autoincrement"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
