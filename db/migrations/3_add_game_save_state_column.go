package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func v3AddGameSaveStateColumn(migrations *migrate.Migrations) error {
	return migrations.Register(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, "ALTER TABLE game_saves ADD COLUMN state text NOT NULL")
		return err
	}, nil)
}
