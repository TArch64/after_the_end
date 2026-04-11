package migrations

import (
	"context"

	"after_the_end/db/model"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func v1AddGameSavesTable(migrations *migrate.Migrations) {
	migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().Model((*model.GameSave)(nil)).Exec(ctx)
		return err
	}, nil)
}
