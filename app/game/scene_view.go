package game

import (
	"after_the_end/backbone"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type SceneView struct {
	*backbone.StatelessView
	model *Model
	hexes map[*model.LocationHex]*SceneHex
}

func NewSceneView(gameModel *Model) *SceneView {
	return &SceneView{
		StatelessView: backbone.NewStatelessView(),
		model:         gameModel,
	}
}

func (v *SceneView) ViewInit() *qt.QWidget {
	v.hexes = make(map[*model.LocationHex]*SceneHex, len(v.model.ActiveLocation.Hexes))
	scene := qt.NewQGraphicsScene()

	for _, locationHex := range v.model.ActiveLocation.Hexes {
		v.hexes[locationHex] = NewSceneHex(scene, locationHex)
	}

	view := qt.NewQGraphicsView3(scene)
	view.SetFrameShape(qt.QFrame__NoFrame)
	return view.QWidget
}
