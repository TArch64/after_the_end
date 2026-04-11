package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func Up(db *bun.DB) (err error) {
	migrations := migrate.NewMigrations()

	v1AddGameSavesTable(migrations)
	v2AddGameSavePosition(migrations)

	ctx := context.Background()
	migrator := migrate.NewMigrator(db, migrations)
	if err = migrator.Init(ctx); err != nil {
		return err
	}

	if _, err = migrator.Migrate(ctx); err != nil {
		return err
	}

	return nil
}
