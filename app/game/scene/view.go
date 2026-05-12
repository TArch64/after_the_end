package scene

import (
	"time"

	"after_the_end/app/game/command"
	"after_the_end/app/game/command/cmd"
	"after_the_end/app/game/state"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db/model"
	"after_the_end/helper/axial"
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
	stateModel       *state.Model
	hexes            map[string]*Hex
	activeHex        *Hex
	graphicsScene    *qt.QGraphicsScene
	graphicsView     *qt.QGraphicsView
	viewportTimeLine *qt.QTimeLine
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
		v.activateHex(v.stateModel.MainCharacter.LocationCoord, false, nil)
	})

	v.AutoDispose(
		command.MainThreadHandle[*cmd.MoveMainCharacter](func(cmd *cmd.MoveMainCharacter) {
			v.activateHex(cmd.Coord, true, cmd.Complete)
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

func (v *View) activateHex(coord *axial.Coord, animated bool, onFinish func()) {
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
	dx := -pos.X() + rect.Width()/2
	dy := -pos.Y() + rect.Height()/2

	endPos := qt.NewQPointF3(dx, dy)
	translate := v.graphicsView.Transform()
	startPos := qt.NewQPointF3(translate.Dx(), translate.Dy())
	delta := endPos.OperatorMinusAssign(startPos)

	if animated {
		v.animateViewTranslate(delta, onFinish)
	} else {
		v.graphicsView.Translate(delta.X(), delta.Y())
	}
}

func (v *View) animateViewTranslate(delta *qt.QPointF, onFinish func()) {
	if v.viewportTimeLine != nil {
		return
	}

	distance := qt.NewQLineF2(qt.NewQPointF(), delta)
	duration := max(500, time.Duration(distance.Length()*1))

	v.viewportTimeLine = qttimer.TimeLine(&qttimer.Translation{
		Duration: duration * time.Millisecond,

		Tick: func(step float64) {
			v.graphicsView.Translate(delta.X()*step, delta.Y()*step)
		},

		Finish: func() {
			v.viewportTimeLine = nil
			onFinish()
		},
	})
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
