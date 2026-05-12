package axial

const NeighbourCount = 6

var neighbours = [NeighbourCount]*Coord{
	{Q: 1, R: 0},
	{Q: -1, R: 0},
	{Q: 0, R: 1},
	{Q: 0, R: -1},
	{Q: 1, R: -1},
	{Q: -1, R: 1},
}

func (c *Coord) Neighbours() [NeighbourCount]*Coord {
	coords := [NeighbourCount]*Coord{}
	for idx, coord := range neighbours {
		coords[idx] = c.Plus(coord.Q, coord.R)
	}
	return coords
}
