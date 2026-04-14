package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func v4AddTimestampIndexes(migrations *migrate.Migrations) error {
	return migrations.Register(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			CREATE INDEX idx_game_saves_created_at ON game_saves(created_at);
			CREATE INDEX idx_game_saves_updated_at ON game_saves(updated_at);
		`)

		return err
	}, nil)
}
