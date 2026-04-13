package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

func v2AddCharactersTable(migrations *migrate.Migrations) error {
	return migrations.Register(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			CREATE TABLE characters (
				id serial NOT NULL PRIMARY KEY,
				type smallint NOT NULL,
				save_id int NOT NULL,
				created_at timestamp NOT NULL DEFAULT current_timestamp,
				updated_at timestamp NOT NULL DEFAULT current_timestamp
			);
		
			CREATE INDEX idx_characters_save_id ON characters(save_id);
			CREATE INDEX idx_characters_type ON characters(type);
		`)

		return err
	}, nil)
}
