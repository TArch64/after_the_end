package gamewizard

import (
	"fmt"

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
		backbone.NewStatefullView(NewModel()),
	}
}

func (v *View) ViewBeforeOpen(params router.Params) error {
	v.Model.Load(params["gameSave"].(*model.GameSave))
	return nil
}

func (v *View) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	_ = qt.NewQLabel5(fmt.Sprintf("wizard for save %d", v.Model.GameSave.ID), widget)
	return widget
}
