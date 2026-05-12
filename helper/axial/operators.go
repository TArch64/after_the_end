package axial

import (
	"after_the_end/helper/mathg"
)

func (c *Coord) Plus(q, r int) *Coord {
	return &Coord{
		Q: c.Q + q,
		R: c.R + r,
	}
}

func (c *Coord) Distance(to *Coord) uint {
	dq := mathg.Abs(c.Q - to.Q)
	dr := mathg.Abs(c.R - to.R)
	ds := mathg.Abs(c.Q + c.R - to.Q - to.R)
	return uint((dq + dr + ds) / 2)
}
