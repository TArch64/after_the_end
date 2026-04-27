package saves

import (
	"after_the_end/app/dialog/confirm"
	"after_the_end/app/resources"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db/model"

	qt "github.com/mappu/miqt/qt6"
)

type SaveAction struct {
	Icon      *qt.QIcon
	OnClicked func()
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
	container.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Fixed)
	container.SetProperty("card", qt.NewQVariant4(2))

	row := qt.NewQHBoxLayout(container)
	row.SetContentsMargins(0, 0, 24, 0)
	row.AddWidget(v.renderInfoColumn())
	row.AddStretch()

	row.AddWidget(v.renderAction(&SaveAction{
		Icon:      resources.QIcon("resume-main"),
		OnClicked: v.resume,
	}))

	row.AddWidget(v.renderAction(&SaveAction{
		Icon:      resources.QIcon("trash-main"),
		OnClicked: v.delete,
	}))

	return container
}

func (v *SaveView) renderInfoColumn() *qt.QWidget {
	widget := qt.NewQWidget2()
	column := qt.NewQVBoxLayout(widget)

	title := qt.NewQLabel3(v.Model.FormatTitle())
	title.SetProperty("text-body", qt.NewQVariant4(1))
	column.AddWidget(title.QWidget)

	column.AddStretch()

	updatedAt := qt.NewQLabel3(v.Model.FormatUpdatedAt())
	title.SetProperty("text-body", qt.NewQVariant4(2))
	column.AddWidget(updatedAt.QWidget)

	return widget
}

func (v *SaveView) renderAction(action *SaveAction) *qt.QWidget {
	button := qt.NewQPushButton2()
	button.SetIcon(action.Icon)
	button.SetIconSize(qt.NewQSize2(32, 32))
	button.SetFixedSize2(40, 40)
	button.SetProperty("button", qt.NewQVariant11("icon-secondary"))
	button.OnClicked(action.OnClicked)
	return button.QWidget
}

func (v *SaveView) resume() {
	if v.Model.GameSave.State != model.GameSaveReady {
		router.Push(router.RouteGameWizard, router.Params{
			"gameSave": v.Model.GameSave,
			"returnTo": router.RouteSaves,
		})
		return
	}

	router.Push(router.RouteGame, router.Params{
		"gameSave": v.Model.GameSave,
	})
}

func (v *SaveView) delete() {
	confirmed := confirm.Show(&confirm.Options{
		Parent: v.ViewRoot(),
		Title:  v.Model.FormatTitle(),
		Text:   "Are you sure you want to delete this save?",
	})

	if confirmed {
		v.onDeleted()
	}
}
