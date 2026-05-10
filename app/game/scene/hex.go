package scene

import (
	"fmt"

	"after_the_end/db/model"

	qt "github.com/mappu/miqt/qt6"
)

const hexSize = 50

type Hex struct {
	gPath       *qt.QGraphicsPathItem
	scene       *qt.QGraphicsScene
	locationHex *model.LocationHex
}

func NewHex(
	scene *qt.QGraphicsScene,
	locationHex *model.LocationHex,
) *Hex {
	item := &Hex{
		scene:       scene,
		locationHex: locationHex,
	}

	item.render()
	return item
}

func (h *Hex) render() {
	h.renderPath()
	h.renderText()
}

func (h *Hex) renderPath() {
	cx := hexSize * 1.5 * float64(h.locationHex.Q)
	cy := hexSize * (sqrt3/2*float64(h.locationHex.Q) + sqrt3*float64(h.locationHex.R))

	h.gPath = qt.NewQGraphicsPathItem2(qt.NewQPainterPath3(hexPath))
	h.gPath.SetPos2(asIso(cx, cy))
	h.gPath.SetBrush(qt.NewQBrush3(qt.NewQColor3(136, 170, 255)))
	h.gPath.SetData(int(KeyHex), qt.NewQVariant14(h.locationHex.StringKey()))
	h.scene.AddItem(h.gPath.QGraphicsItem)
}

func (h *Hex) renderText() {
	text := qt.NewQGraphicsTextItem2(h.locationHex.StringKey())
	h.addChild(text.QGraphicsItem)

	rect := text.BoundingRect()
	text.SetPos2(-rect.Width()/2, -rect.Height()/2)
}

func (h *Hex) addChild(child *qt.QGraphicsItem) {
	child.SetParentItem(h.gPath.QGraphicsItem)
	child.SetAcceptedMouseButtons(qt.NoButton)
}

func (h *Hex) OnClicked() {
	fmt.Printf("clicked on %s\n", h.locationHex.StringKey())
}

func (h *Hex) Delete() {
	h.gPath.Delete()
}
