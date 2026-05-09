package scene

import (
	qt "github.com/mappu/miqt/qt6"
)

const (
	sqrt3 = 1.7320508075688772 // math.Sqrt(3)
	cos30 = 0.8660254037844386 // math.Cos(math.Pi/6)
	sin30 = 0.5                // math.Sin(math.Pi/6)
)

var (
	hexPath = newHexPath()
)

func newHexPath() *qt.QPainterPath {
	const (
		hexSideX = 0.5 * hexSize
		hexSideY = (sqrt3 / 2) * hexSize
	)

	path := qt.NewQPainterPath()

	path.MoveTo2(asIso(hexSize, 0))
	path.LineTo2(asIso(hexSideX, hexSideY))
	path.LineTo2(asIso(-hexSideX, hexSideY))
	path.LineTo2(asIso(-hexSize, 0))
	path.LineTo2(asIso(-hexSideX, -hexSideY))
	path.LineTo2(asIso(hexSideX, -hexSideY))

	path.CloseSubpath()
	return path
}

func asIso(x, y float64) (float64, float64) {
	return (x - y) * cos30, (x + y) * sin30
}
