package mapgenerator

import (
	"context"
	"fmt"

	"after_the_end/db"
	"after_the_end/db/model"
	"after_the_end/helper/axial"

	"github.com/uptrace/bun"
)

const (
	height = 20
	width  = 30
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

	for coord := range axial.RectIterator(width, height) {
		location.Hexes = append(location.Hexes, &model.LocationHex{
			Coord: coord,
		})
	}

	return location
}
