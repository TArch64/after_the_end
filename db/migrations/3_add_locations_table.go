package migrations

import (
	"github.com/uptrace/bun/migrate"
)

func v3AddLocationsTable(migrations *migrate.Migrations) error {
	return migrations.Register(ddl(`
		CREATE TABLE locations (
			id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			name text NOT NULL,
			save_id int NOT NULL REFERENCES game_saves (id) ON DELETE CASCADE
		);

		CREATE INDEX idx_locations_save_id ON locations(save_id);

		CREATE TABLE location_hexes (
			q int NOT NULL,
			r int NOT NULL,
			elevation int NOT NULL,
			coord text NOT NULL GENERATED ALWAYS AS (CAST(q AS text) || ':' || CAST(r AS text)) STORED,
			location_id int NOT NULL REFERENCES locations (id) ON DELETE CASCADE,
			PRIMARY KEY (location_id, q, r)
		);

		CREATE INDEX idx_location_hexes_location_id ON location_hexes(location_id);
		CREATE UNIQUE INDEX idx_location_hexes_coord ON location_hexes (location_id, coord);

		ALTER TABLE characters ADD COLUMN location_id int;
		ALTER TABLE characters ADD COLUMN location_coord text NOT NULL DEFAULT '0:0';
	`))
}
