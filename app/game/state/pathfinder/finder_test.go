package pathfinder

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"after_the_end/db/model"
	"after_the_end/helper/axial"
)

func newLocation(coords ...*axial.Coord) *model.Location {
	hexes := make([]*model.LocationHex, len(coords))
	for idx, coord := range coords {
		hexes[idx] = &model.LocationHex{Coord: coord}
	}

	return &model.Location{Hexes: hexes}
}

type CoordSet []*axial.Coord

func (s CoordSet) StringList() []string {
	coords := make([]string, len(s))
	for idx, coord := range s {
		coords[idx] = coord.StringKey()
	}
	return coords
}

func TestFinder_Find(t *testing.T) {
	location := newLocation(
		&axial.Coord{Q: 0, R: 0},
		&axial.Coord{Q: -1, R: 0},
		&axial.Coord{Q: 0, R: -1},
		&axial.Coord{Q: 1, R: -1},
		&axial.Coord{Q: 0, R: 1},
		&axial.Coord{Q: -1, R: 1},
		&axial.Coord{Q: 1, R: 1},
		&axial.Coord{Q: 2, R: 0},
		&axial.Coord{Q: 1, R: 2},
		&axial.Coord{Q: 2, R: 1},
	)

	type TestCase struct {
		Name   string
		From   *axial.Coord
		To     *axial.Coord
		Expect CoordSet
	}

	testCases := []*TestCase{
		{
			Name: "find path",
			From: &axial.Coord{Q: 0, R: 0},
			To:   &axial.Coord{Q: 2, R: 1},

			Expect: CoordSet{
				&axial.Coord{Q: 0, R: 0},
				&axial.Coord{Q: 0, R: 1},
				&axial.Coord{Q: 1, R: 1},
				&axial.Coord{Q: 2, R: 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			finder := New(location, tc.From, tc.To)
			expect := tc.Expect.StringList()
			actual := CoordSet(finder.Find()).StringList()

			if diff := cmp.Diff(expect, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
