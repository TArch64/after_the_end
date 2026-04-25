package game

import (
	"math"

	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

const hexSize = 50

var (
	sqrt3 = math.Sqrt(3)

	hexCorners = [6]*qt.QPointF{
		qt.NewQPointF3(1, 0),
		qt.NewQPointF3(0.5, sqrt3/2),
		qt.NewQPointF3(-0.5, sqrt3/2),
		qt.NewQPointF3(-1, 0),
		qt.NewQPointF3(-0.5, -sqrt3/2),
		qt.NewQPointF3(0.5, -sqrt3/2),
	}
)

type SceneHex struct {
	*qt.QGraphicsPathItem
	scene       *qt.QGraphicsScene
	locationHex *model.LocationHex
}

func NewSceneHex(
	scene *qt.QGraphicsScene,
	locationHex *model.LocationHex,
) *SceneHex {
	item := &SceneHex{
		QGraphicsPathItem: qt.NewQGraphicsPathItem(),
		scene:             scene,
		locationHex:       locationHex,
	}

	item.render()
	return item
}

func (h *SceneHex) render() {
	path := h.renderPath()
	h.QGraphicsPathItem = h.scene.AddPath(path)
	h.QGraphicsPathItem.SetBrush(qt.NewQBrush3(qt.NewQColor3(136, 170, 255)))
}

func (h *SceneHex) renderPath() *qt.QPainterPath {
	cx := hexSize * 1.5 * float64(h.locationHex.Q)
	cy := hexSize * (sqrt3/2*float64(h.locationHex.Q) + sqrt3*float64(h.locationHex.R))

	path := qt.NewQPainterPath()
	path.MoveTo2(cx+hexCorners[0].X()*hexSize, cy+hexCorners[0].Y()*hexSize)
	for idx := 1; idx < len(hexCorners); idx++ {
		path.LineTo2(cx+hexCorners[idx].X()*hexSize, cy+hexCorners[idx].Y()*hexSize)
	}

	path.CloseSubpath()
	return path
}
