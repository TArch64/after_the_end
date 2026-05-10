package migrations

import (
	"github.com/uptrace/bun/migrate"
)

func v1AddGameSavesTable(migrations *migrate.Migrations) error {
	return migrations.Register(ddl(`
		CREATE TABLE game_saves (
			id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			position smallint NOT NULL,
			state text NOT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
			
		CREATE INDEX idx_game_saves_created_at ON game_saves(created_at);
		CREATE INDEX idx_game_saves_updated_at ON game_saves(updated_at);
	`))
}
