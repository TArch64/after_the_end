package scene

import (
	"after_the_end/app/game/state"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db/model"

	qt "github.com/mappu/miqt/qt6"
)

const (
	guiWidth  = 20
	guiHeight = 10
)

type View struct {
	*backbone.StatelessView
	stateModel    *state.Model
	hexes         map[string]*Hex
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

	v.hexes = make(map[string]*Hex, len(location.Hexes))
	for _, locationHex := range location.Hexes {
		v.hexes[locationHex.StringKey()] = NewHex(v.graphicsScene, locationHex)
	}
}

func (v *View) renderGraphicsView() {
	v.graphicsView = qt.NewQGraphicsView3(v.graphicsScene)
	v.panning.View = v.graphicsView

	v.graphicsView.SetFrameShape(qt.QFrame__NoFrame)
	v.graphicsView.SetVerticalScrollBarPolicy(qt.ScrollBarAlwaysOff)
	v.graphicsView.SetHorizontalScrollBarPolicy(qt.ScrollBarAlwaysOff)

	v.graphicsView.OnWheelEvent(v.onWheelEvent)
	v.graphicsView.OnMousePressEvent(v.onMousePressEvent)
	v.graphicsView.OnMouseMoveEvent(v.onMouseMoveEvent)
	v.graphicsView.OnMouseReleaseEvent(v.onMouseReleaseEvent)
	v.graphicsView.OnResizeEvent(v.onResizeEvent)
}

func (v *View) onWheelEvent(_ func(event *qt.QWheelEvent), _ *qt.QWheelEvent) {}

func (v *View) onMousePressEvent(_ func(event *qt.QMouseEvent), event *qt.QMouseEvent) {
	v.panning.Start(event)
}

func (v *View) onMouseMoveEvent(_ func(event *qt.QMouseEvent), event *qt.QMouseEvent) {
	v.panning.Move(event)
}

func (v *View) onMouseReleaseEvent(_ func(event *qt.QMouseEvent), event *qt.QMouseEvent) {
	if v.panning.End() {
		return
	}

	item := v.graphicsView.ItemAt(event.Pos())
	if item == nil {
		return
	}

	if key := item.Data(int(KeyHex)).ToString(); key != "" {
		v.hexes[key].OnClicked()
	}
}

func (v *View) onResizeEvent(super func(event *qt.QResizeEvent), event *qt.QResizeEvent) {
	super(event)
	boundingRect := v.graphicsScene.ItemsBoundingRect()
	padX := float64(v.graphicsView.Viewport().Width()) / 4
	padY := float64(v.graphicsView.Viewport().Height()) / 4
	v.graphicsView.SetSceneRect(boundingRect.Adjusted(-padX, -padY, padX, padY))
}

func (v *View) renderBackButton() *qt.QWidget {
	button := qt.NewQPushButton3("back")
	button.SetContentsMargins(0, 0, 0, 0)
	button.SetProperty("button", qt.NewQVariant11("main"))

	button.OnClicked(func() {
		router.Push(router.RouteStart)
	})

	return button.QWidget
}
