package game

import (
	"after_the_end/backbone"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type SceneView struct {
	*backbone.StatelessView
	model         *Model
	hexes         map[*model.LocationHex]*SceneHex
	graphicsScene *qt.QGraphicsScene
	graphicsView  *qt.QGraphicsView
	panning       *ScenePanning
}

func NewSceneView(gameModel *Model) *SceneView {
	return &SceneView{
		StatelessView: backbone.NewStatelessView(),
		model:         gameModel,
		panning:       NewScenePanning(),
	}
}

func (v *SceneView) ViewInit() *qt.QWidget {
	v.graphicsScene = qt.NewQGraphicsScene()
	v.renderLocation(v.model.ActiveLocation)
	return v.renderGraphicsView()
}

func (v *SceneView) renderLocation(location *model.Location) {
	if v.hexes != nil {
		for _, hex := range v.hexes {
			hex.Delete()
		}
	}

	v.hexes = make(map[*model.LocationHex]*SceneHex, len(location.Hexes))
	for _, locationHex := range location.Hexes {
		v.hexes[locationHex] = NewSceneHex(v.graphicsScene, locationHex)
	}
}

func (v *SceneView) renderGraphicsView() *qt.QWidget {
	v.graphicsView = qt.NewQGraphicsView3(v.graphicsScene)
	v.panning.View = v.graphicsView

	v.graphicsView.SetFrameShape(qt.QFrame__NoFrame)

	v.graphicsView.SetVerticalScrollBarPolicy(qt.ScrollBarAlwaysOff)
	v.graphicsView.SetHorizontalScrollBarPolicy(qt.ScrollBarAlwaysOff)
	v.graphicsView.OnWheelEvent(func(super func(event *qt.QWheelEvent), event *qt.QWheelEvent) {})

	v.graphicsView.OnMousePressEvent(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent) {
		v.panning.Start(event)
	})

	v.graphicsView.OnMouseMoveEvent(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent) {
		v.panning.Move(event)
	})

	v.graphicsView.OnMouseReleaseEvent(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent) {
		v.panning.End()
	})

	v.graphicsView.OnResizeEvent(func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent) {
		super(event)
		v.resizeView()
	})

	return v.graphicsView.QWidget
}

func (v *SceneView) resizeView() {
	boundingRect := v.graphicsScene.ItemsBoundingRect()
	padX := float64(v.graphicsView.Viewport().Width()) / 4
	padY := float64(v.graphicsView.Viewport().Height()) / 4
	v.graphicsView.SetSceneRect(boundingRect.Adjusted(-padX, -padY, padX, padY))
}
