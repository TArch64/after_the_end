package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Character struct {
	bun.BaseModel `bun:"table:characters,alias:c"`
	ID            ID            `bun:",pk,autoincrement"`
	Type          CharacterType `bun:",notnull"`
	SaveID        ID            `bun:",notnull"`
	CreatedAt     time.Time     `bun:",notnull,default:current_timestamp"`
	UpdatedAt     time.Time     `bun:",notnull,default:current_timestamp"`

	// Relations
	Save *GameSave `bun:"rel:belongs-to,join:save_id=id"`
}
