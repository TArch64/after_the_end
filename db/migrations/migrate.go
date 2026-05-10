package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var migrationFuncs = []func(migrations *migrate.Migrations) error{
	v1AddGameSavesTable,
	v2AddLocationsTable,
	v3AddCharactersTable,
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

func ddl(query string) (migrate.MigrationFunc, migrate.MigrationFunc) {
	return func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, query)
		return err
	}, nil
}
