package game

import (
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type View struct {
	*backbone.StatefullView[*Model]
}

func NewView() *View {
	return &View{
		StatefullView: backbone.NewStatefullView(NewModel()),
	}
}

func (v *View) ViewBeforeOpen(params router.Params) error {
	return v.Model.Load(params["gameSave"].(*model.GameSave))
}

func (v *View) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	column := qt.NewQStackedLayout(widget)
	column.AddWidget(v.Mount(NewSceneView(v.Model)))
	return widget
}
