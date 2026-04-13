package saves

import (
	"after_the_end/app/dialog/confirm"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type SaveAction struct {
	Name      string
	Icon      *qt.QIcon
	OnPressed func()
}

type SaveView struct {
	*backbone.StatefullView[*SaveModel]
	onDeleted func()
}

type SaveViewOptions struct {
	SavesModel *Model
	GameSave   *model.GameSave
	OnDeleted  func()
}

func NewSaveView(options *SaveViewOptions) *SaveView {
	return &SaveView{
		StatefullView: backbone.NewStatefullView(
			NewSaveModel(options.SavesModel, options.GameSave),
		),
		onDeleted: options.OnDeleted,
	}
}

func (v *SaveView) ViewInit() *qt.QWidget {
	container := qt.NewQWidget2()
	container.SetObjectName("saves_list_item")
	container.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Fixed)
	container.SetStyleSheet(styled.S("#saves_list_item", styled.Card2))

	row := qt.NewQHBoxLayout(container)
	row.SetContentsMargins(0, 0, 24, 0)
	row.AddWidget(v.renderInfoColumn())
	row.AddStretch()

	row.AddWidget(v.renderAction(&SaveAction{
		Name:      "save_delete",
		Icon:      qt.NewQIcon4(":/icons/trash-main.svg"),
		OnPressed: v.delete,
	}))

	return container
}

func (v *SaveView) renderInfoColumn() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("save_info")

	column := qt.NewQVBoxLayout(widget)
	column.SetObjectName("save_info")

	title := qt.NewQLabel3(v.Model.FormatTitle())
	title.SetObjectName("save_title")
	title.SetStyleSheet(styled.Body)
	column.AddWidget(title.QWidget)

	column.AddStretch()

	updatedAt := qt.NewQLabel3(v.Model.FormatUpdatedAt())
	updatedAt.SetObjectName("save_updated_at")
	updatedAt.SetStyleSheet(styled.Body2)
	column.AddWidget(updatedAt.QWidget)

	return widget
}

func (v *SaveView) renderAction(action *SaveAction) *qt.QWidget {
	button := qt.NewQPushButton2()
	button.SetObjectName(action.Name)
	button.SetIcon(action.Icon)
	button.SetIconSize(qt.NewQSize2(32, 32))
	button.SetFixedSize2(40, 40)
	button.SetStyleSheet(styled.ButtonIconSecondary)
	button.OnReleased(action.OnPressed)
	return button.QWidget
}

func (v *SaveView) delete() {
	confirmed := confirm.Show(&confirm.Options{
		Parent: v.Root,
		Title:  v.Model.FormatTitle(),
		Text:   "Are you sure you want to delete this save?",
	})

	if confirmed {
		v.onDeleted()
	}
}
