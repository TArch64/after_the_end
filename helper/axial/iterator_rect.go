package axial

import (
	"iter"

	"after_the_end/helper/mathg"
)

func RectIterator(width, height int) iter.Seq[*Coord] {
	return func(yield func(*Coord) bool) {
		radius := max(width, height)

		for q := -radius; q <= radius; q++ {
			for r := -radius; r <= radius; r++ {
				s := -q - r
				if mathg.Abs(r-s) >= height {
					continue
				}
				if mathg.Abs(q)*2 > width {
					continue
				}
				yield(&Coord{
					Q: q,
					R: r,
					S: s,
				})
			}
		}
	}
}
