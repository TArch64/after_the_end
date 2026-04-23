package game

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
		StatefullView: backbone.NewStatefullView(NewModel()),
	}
}

func (v *View) ViewBeforeOpen(params router.Params) error {
	v.Model.Load(params["gameSave"].(*model.GameSave))
	return nil
}

func (v *View) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	column := qt.NewQVBoxLayout(widget)

	text := qt.NewQLabel3(fmt.Sprintf("save #%d", v.Model.GameSave.ID))
	column.AddWidget(text.QWidget)

	return widget
}
