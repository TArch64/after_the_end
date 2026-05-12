package scene

import (
	"after_the_end/helper/axial"

	qt "github.com/mappu/miqt/qt6"
)

const (
	sqrt3   = 1.7320508075688772 // math.Sqrt(3)
	cos30   = 0.8660254037844386 // math.Cos(math.Pi/6)
	sin30   = 0.5                // math.Sin(math.Pi/6)
	hexSize = 50
)

var (
	hexPath        = newHexPath(0)
	hexOutlinePath = newHexPath(2)
)

func newHexPath(inset float64) *qt.QPainterPath {
	var (
		size     = hexSize - inset
		hexSideX = 0.5 * size
		hexSideY = (sqrt3 / 2) * size
	)

	path := qt.NewQPainterPath()

	path.MoveTo2(asIso(size, 0))
	path.LineTo2(asIso(hexSideX, hexSideY))
	path.LineTo2(asIso(-hexSideX, hexSideY))
	path.LineTo2(asIso(-size, 0))
	path.LineTo2(asIso(-hexSideX, -hexSideY))
	path.LineTo2(asIso(hexSideX, -hexSideY))

	path.CloseSubpath()
	return path
}

func HexCenterPos(coord *axial.Coord) (float64, float64) {
	cx := hexSize * 1.5 * float64(coord.Q)
	cy := hexSize * (sqrt3/2*float64(coord.Q) + sqrt3*float64(coord.R))
	return asIso(cx, cy)
}

func asIso(x, y float64) (float64, float64) {
	return (x - y) * cos30, (x + y) * sin30
}
