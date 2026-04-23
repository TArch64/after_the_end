package mapgenerator

import (
	"context"
	"fmt"

	"after_the_end/db"
	"after_the_end/db/model"
	"after_the_end/helper/mathg"

	"github.com/uptrace/bun"
)

func Generate(ctx context.Context, gameSave *model.GameSave) error {
	worldMap := generateWorldMap()

	return db.Tx(func(tx bun.Tx) error {
		worldMap.SaveID = gameSave.ID

		_, err := tx.NewInsert().
			Model(worldMap).
			Exec(ctx)

		if err != nil {
			return fmt.Errorf("insert location: %w", err)
		}

		for _, hex := range worldMap.Hexes {
			hex.LocationID = worldMap.ID
		}

		_, err = tx.NewInsert().
			Model(&worldMap.Hexes).
			Exec(ctx)

		if err != nil {
			return fmt.Errorf("insert location hexes: %w", err)
		}

		return nil
	})
}

func generateWorldMap() *model.Location {
	location := &model.Location{
		Name: "World Map",
	}

	const radius = 4

	for q := -radius; q <= radius; q++ {
		for r := -radius; r <= radius; r++ {
			s := -q - r
			if mathg.Abs(s) <= radius {
				location.Hexes = append(location.Hexes, &model.LocationHex{
					Q: q,
					R: r,
					S: s,
				})
			}
		}
	}

	return location
}
