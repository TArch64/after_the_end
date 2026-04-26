package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var migrationFuncs = []func(migrations *migrate.Migrations) error{
	v1AddGameSavesTable,
	v2AddCharactersTable,
	v3AddLocationsTable,
}

func Up(db *bun.DB) (err error) {
	migrations := migrate.NewMigrations()
	for _, register := range migrationFuncs {
		if err = register(migrations); err != nil {
			return err
		}
	}
	ctx := context.Background()
	migrator := migrate.NewMigrator(db, migrations,
		migrate.WithMarkAppliedOnSuccess(true),
	)
	if err = migrator.Init(ctx); err != nil {
		return err
	}
	if _, err = migrator.Migrate(ctx); err != nil {
		return err
	}
	return nil
}
