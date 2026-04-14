package start

import (
	"after_the_end/app/dialog/errorreport"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
)

type MenuView struct {
	*backbone.StatelessView
	model *Model
}

type MenuItem struct {
	Title     string
	OnPressed func()
}

func NewMenuView(model *Model) *MenuView {
	return &MenuView{
		StatelessView: backbone.NewStatelessView(),
		model:         model,
	}
}

func (v *MenuView) ViewInit() *qt.QWidget {
	container := qt.NewQWidget2()
	container.SetObjectName("start_menu")

	layout := qt.NewQVBoxLayout(container)
	layout.SetContentsMargins(0, 50, 0, 0)
	layout.SetObjectName("start_menu")

	layout.AddWidget(v.renderMenuItem(&MenuItem{
		Title:     "New Game",
		OnPressed: v.createNewGame,
	}))

	if v.model.SavesCount != 0 {
		layout.AddWidget(v.renderMenuItem(&MenuItem{
			Title: "Load Game",

			OnPressed: func() {
				router.Push(router.RouteSaves)
			},
		}))
	}

	layout.AddWidget(v.renderMenuItem(&MenuItem{
		Title:     "Exit",
		OnPressed: qt.QCoreApplication_Quit,
	}))

	return container
}

func (v *MenuView) renderMenuItem(item *MenuItem) *qt.QWidget {
	button := qt.NewQPushButton3(item.Title)
	button.OnReleased(item.OnPressed)
	button.SetStyleSheet(styled.Button)
	return button.QWidget
}

func (v *MenuView) createNewGame() {
	gameSave, err := v.model.NewGame()
	if err != nil {
		errorreport.Show(v.Root, err)
	}

	router.Push(router.RouteGameWizard, router.Params{
		"gameSave": gameSave,
	})
}
