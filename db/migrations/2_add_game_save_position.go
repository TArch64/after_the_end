package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func v2AddGameSavePosition(migrations *migrate.Migrations) {
	migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, "ALTER TABLE game_saves ADD COLUMN position int NOT NULL")
		return err
	}, nil)
}
