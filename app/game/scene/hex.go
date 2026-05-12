package scene

import (
	"after_the_end/app/game/command"
	"after_the_end/app/game/command/cmd"
	"after_the_end/db/model"

	qt "github.com/mappu/miqt/qt6"
)

type Hex struct {
	root        *qt.QGraphicsPathItem
	outline     *qt.QGraphicsPathItem
	scene       *qt.QGraphicsScene
	locationHex *model.LocationHex
	Active      bool
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

func (h *Hex) Item() *qt.QGraphicsItem {
	return h.root.QGraphicsItem
}

func (h *Hex) render() {
	h.renderPath()
	h.renderText()
}

func (h *Hex) renderPath() {
	h.root = qt.NewQGraphicsPathItem2(qt.NewQPainterPath3(hexPath))
	h.root.SetPos2(HexCenterPos(h.locationHex.Coord))
	h.root.SetPen(qt.NewQPen2(qt.NoPen))
	h.root.SetBrush(qt.NewQBrush3(qt.NewQColor3(136, 170, 255)))
	h.scene.AddItem(h.root.QGraphicsItem)

	h.root.OnSceneEventFilter(func(_ func(_ *qt.QGraphicsItem, _ *qt.QEvent) bool, _ *qt.QGraphicsItem, _ *qt.QEvent) bool {
		return true
	})

	h.root.OnMousePressEvent(func(_ func(event *qt.QGraphicsSceneMouseEvent), event *qt.QGraphicsSceneMouseEvent) {
		event.Accept()
	})

	h.root.OnMouseReleaseEvent(h.OnClicked)
}

func (h *Hex) renderText() {
	text := qt.NewQGraphicsTextItem2(h.locationHex.Coord.StringKey())
	rect := text.BoundingRect()
	text.SetPos2(-rect.Width()/2, -rect.Height()/2)

	h.addChild(text.QGraphicsItem)
}

func (h *Hex) addChild(child *qt.QGraphicsItem) {
	child.SetParentItem(h.root.QGraphicsItem)
	child.SetAcceptedMouseButtons(qt.NoButton)
	child.InstallSceneEventFilter(h.root.QGraphicsItem)
}

func (h *Hex) OnClicked(_ func(event *qt.QGraphicsSceneMouseEvent), _ *qt.QGraphicsSceneMouseEvent) {
	if h.Active {
		return
	}
	command.Dispatch(cmd.NewWalkPath(h.locationHex.Coord))
}

func (h *Hex) SetActive() {
	h.outline = qt.NewQGraphicsPathItem2(qt.NewQPainterPath3(hexOutlinePath))
	brush := qt.NewQBrush3(qt.NewQColor3(180, 136, 255))
	h.outline.SetPen(qt.NewQPen4(brush, 4))

	h.addChild(h.outline.QGraphicsItem)
	h.Active = true
}

func (h *Hex) SetInactive() {
	h.outline.Delete()
	h.outline = nil
	h.Active = false
}

func (h *Hex) Delete() {
	h.root.Delete()
}
