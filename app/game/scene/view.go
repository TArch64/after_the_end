package scene

import (
	"after_the_end/app/game/state"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

const (
	guiWidth  = 20
	guiHeight = 10
)

type View struct {
	*backbone.StatelessView
	stateModel    *state.Model
	hexes         map[*model.LocationHex]*Hex
	graphicsScene *qt.QGraphicsScene
	graphicsView  *qt.QGraphicsView
	panning       *Panning
}

func NewView(stateModel *state.Model) *View {
	return &View{
		StatelessView: backbone.NewStatelessView(),
		stateModel:    stateModel,
		panning:       NewPanning(),
	}
}

func (v *View) ViewInit() *qt.QWidget {
	v.graphicsScene = qt.NewQGraphicsScene()
	v.renderLocation(v.stateModel.ActiveLocation)
	v.renderGraphicsView()

	widget := qt.NewQWidget2()
	grid := qt.NewQGridLayout(widget)
	grid.SetContentsMargins(0, 0, 0, 0)
	grid.SetSpacing(0)
	grid.AddWidget3(v.graphicsView.QWidget, 0, 0, guiHeight, guiWidth)
	grid.AddWidget2(v.renderBackButton(), 0, guiWidth-1)
	return widget
}

func (v *View) renderLocation(location *model.Location) {
	if v.hexes != nil {
		for _, hex := range v.hexes {
			hex.Delete()
		}
	}

	v.hexes = make(map[*model.LocationHex]*Hex, len(location.Hexes))
	for _, locationHex := range location.Hexes {
		v.hexes[locationHex] = NewHex(v.graphicsScene, locationHex)
	}
}

func (v *View) renderGraphicsView() {
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
}

func (v *View) resizeView() {
	boundingRect := v.graphicsScene.ItemsBoundingRect()
	padX := float64(v.graphicsView.Viewport().Width()) / 4
	padY := float64(v.graphicsView.Viewport().Height()) / 4
	v.graphicsView.SetSceneRect(boundingRect.Adjusted(-padX, -padY, padX, padY))
}

func (v *View) renderBackButton() *qt.QWidget {
	button := qt.NewQPushButton3("back")
	button.SetContentsMargins(0, 0, 0, 0)
	button.SetObjectName("gui_back")
	button.SetStyleSheet(styled.Button + "#gui_back { margin: 0 }")

	button.OnClicked(func() {
		router.Push(router.RouteStart)
	})

	return button.QWidget
}
