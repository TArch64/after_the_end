package saves

import (
	"fmt"

	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type SaveView struct {
	*backbone.BaseView
	gameSave *model.GameSave
}

func NewSaveView(save *model.GameSave) *SaveView {
	return &SaveView{
		BaseView: backbone.NewBaseView(),
		gameSave: save,
	}
}

func (v *SaveView) ViewInit(parent *qt.QWidget) {
	parent.SetStyleSheet(styled.ListItem)
	column := qt.NewQVBoxLayout(parent)

	title := qt.NewQLabel3(fmt.Sprintf("Save #%d", v.gameSave.ID))
	title.SetStyleSheet(styled.BodyWhite)
	column.AddWidget(title.QWidget)
}
