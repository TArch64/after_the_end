package scene

import (
	"after_the_end/app/game/command"
	"after_the_end/app/game/command/cmd"
	"after_the_end/app/game/state"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db/model"
	"after_the_end/helper/qttimer"

	qt "github.com/mappu/miqt/qt6"
	"github.com/mappu/miqt/qt6/opengl"
)

const (
	guiWidth  = 20
	guiHeight = 10
)

type View struct {
	*backbone.StatelessView
	stateModel    *state.Model
	hexes         map[string]*Hex
	activeHex     *Hex
	graphicsScene *qt.QGraphicsScene
	graphicsView  *qt.QGraphicsView
}

func NewView(stateModel *state.Model) *View {
	return &View{
		StatelessView: backbone.NewStatelessView(),
		stateModel:    stateModel,
	}
}

func (v *View) ViewInit() *qt.QWidget {
	v.graphicsScene = qt.NewQGraphicsScene()
	v.renderLocation(v.stateModel.ActiveLocation)
	v.renderGraphicsView()

	qttimer.NextTick(func() {
		v.activateHex(v.stateModel.MainCharacter.LocationCoord)
	})

	v.AutoDispose(
		command.MainThreadHandle[*cmd.ActivateHex](func(cmd *cmd.ActivateHex) {
			v.activateHex(cmd.Coord)
		}),
	)

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
		v.hexes[locationHex.Coord.StringKey()] = NewHex(v.graphicsScene, locationHex)
	}
}

func (v *View) renderGraphicsView() {
	v.graphicsView = qt.NewQGraphicsView3(v.graphicsScene)

	v.graphicsView.SetViewport(opengl.NewQOpenGLWidget2().QWidget)
	v.graphicsView.SetFrameShape(qt.QFrame__NoFrame)

	v.graphicsView.SetVerticalScrollBarPolicy(qt.ScrollBarAlwaysOff)
	v.graphicsView.SetHorizontalScrollBarPolicy(qt.ScrollBarAlwaysOff)
	v.graphicsView.OnWheelEvent(func(_ func(event *qt.QWheelEvent), _ *qt.QWheelEvent) {})
	v.graphicsView.OnScrollContentsBy(func(_ func(dx int, dy int), _ int, _ int) {})
}

func (v *View) activateHex(coord *model.AxialCoord) {
	activatingHex := v.hexes[coord.StringKey()]
	if activatingHex.Active {
		return
	}

	if v.activeHex != nil {
		v.activeHex.SetInactive()
	}

	v.activeHex = activatingHex
	v.activeHex.SetActive()

	pos := v.activeHex.Item().Pos()
	rect := v.graphicsView.Rect().ToRectF()
	translate := v.graphicsView.Transform()
	dx := -pos.X() + rect.Width()/2 - translate.Dx()
	dy := -pos.Y() + rect.Height()/2 - translate.Dy()
	v.graphicsView.Translate(dx, dy)
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
