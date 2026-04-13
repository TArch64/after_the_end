package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func v3AddGameSaveState(migrations *migrate.Migrations) error {
	return migrations.Register(func(ctx context.Context, db *bun.DB) error {

	}, nil)
}
