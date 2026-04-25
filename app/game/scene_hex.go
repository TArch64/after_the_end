package game

import (
	"fmt"
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
	gPath       *qt.QGraphicsPathItem
	gText       *qt.QGraphicsTextItem
	scene       *qt.QGraphicsScene
	locationHex *model.LocationHex
}

func NewSceneHex(
	scene *qt.QGraphicsScene,
	locationHex *model.LocationHex,
) *SceneHex {
	item := &SceneHex{
		scene:       scene,
		locationHex: locationHex,
	}

	item.render()
	return item
}

func (h *SceneHex) render() {
	h.renderPath()
	h.renderText()
}

func (h *SceneHex) renderPath() {
	cx := hexSize * 1.5 * float64(h.locationHex.Q)
	cy := hexSize * (sqrt3/2*float64(h.locationHex.Q) + sqrt3*float64(h.locationHex.R))

	path := qt.NewQPainterPath()
	path.MoveTo2(hexCorners[0].X()*hexSize, hexCorners[0].Y()*hexSize)
	for idx := 1; idx < len(hexCorners); idx++ {
		path.LineTo2(hexCorners[idx].X()*hexSize, hexCorners[idx].Y()*hexSize)
	}

	path.CloseSubpath()
	h.gPath = h.scene.AddPath(path)
	h.gPath.SetPos2(cx, cy)
	h.gPath.SetBrush(qt.NewQBrush3(qt.NewQColor3(136, 170, 255)))
}

func (h *SceneHex) renderText() {
	h.gText = qt.NewQGraphicsTextItem4(
		fmt.Sprintf("q %d\nr %d\ns %d",
			h.locationHex.Q,
			h.locationHex.R,
			h.locationHex.S,
		),
		h.gPath.QGraphicsItem,
	)

	rect := h.gText.BoundingRect()
	h.gText.SetPos2(-rect.Width()/2, -rect.Height()/2)
}

func (h *SceneHex) Delete() {
	h.gPath.Delete()
}
