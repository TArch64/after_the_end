package axial

import (
	"iter"

	"after_the_end/helper/mathg"
)

func RectSeq(width, height int) iter.Seq[*Coord] {
	return func(yield func(*Coord) bool) {
		radius := max(width, height)

		for q := -radius; q <= radius; q++ {
			for r := -radius; r <= radius; r++ {
				if mathg.Abs(2*r+q) >= height {
					continue
				}
				if mathg.Abs(2*q) > width {
					continue
				}
				yield(&Coord{
					Q: q,
					R: r,
				})
			}
		}
	}
}
