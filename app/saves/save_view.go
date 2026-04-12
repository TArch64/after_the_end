package saves

import (
	"fmt"

	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type SaveView struct {
	*backbone.StatefullView[*SaveModel]
}

func NewSaveView(save *model.GameSave) *SaveView {
	return &SaveView{
		StatefullView: backbone.NewStatefullView(NewSaveModel(save)),
	}
}

func (v *SaveView) ViewInit(parent *qt.QWidget) {
	parent.SetObjectName("saves_list_item")
	parent.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Fixed)
	parent.SetStyleSheet(styled.S("#saves_list_item", styled.Card2))

	row := qt.NewQHBoxLayout(parent)
	row.AddWidget(v.renderInfoColumn())
}

func (v *SaveView) renderInfoColumn() *qt.QWidget {
	widget := qt.NewQWidget2()
	column := qt.NewQVBoxLayout(widget)

	title := qt.NewQLabel3(fmt.Sprintf("Save #%d", v.Model.GameSave.ID))
	title.SetStyleSheet(styled.BodyWhite)
	column.AddWidget(title.QWidget)

	column.AddStretch()

	updatedAt := qt.NewQLabel3(v.Model.FormatUpdatedAt())
	updatedAt.SetStyleSheet(styled.BodyWhite2)
	column.AddWidget(updatedAt.QWidget)

	return widget
}
