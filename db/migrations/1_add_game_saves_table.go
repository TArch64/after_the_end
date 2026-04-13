package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func v1AddGameSavesTable(migrations *migrate.Migrations) error {
	return migrations.Register(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			CREATE TABLE game_saves (
				id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
				position smallint NOT NULL,
				created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
			)
		`)

		return err
	}, nil)
}
