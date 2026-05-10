package migrations

import (
	"github.com/uptrace/bun/migrate"
)

func v2AddLocationsTable(migrations *migrate.Migrations) error {
	return migrations.Register(ddl(`
		CREATE TABLE locations (
			id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			name text NOT NULL,
			save_id int NOT NULL REFERENCES game_saves (id) ON DELETE CASCADE
		);

		CREATE INDEX idx_locations_save_id ON locations(save_id);

		CREATE TABLE location_hexes (
			coord text NOT NULL,
			elevation int NOT NULL,
			location_id int NOT NULL REFERENCES locations (id) ON DELETE CASCADE,
			PRIMARY KEY (location_id, coord)
		);

		CREATE INDEX idx_location_hexes_location_id ON location_hexes (location_id);
	`))
}
