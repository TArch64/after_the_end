package start

import (
	"after_the_end/app/dialog/reporterr"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/db/model"

	"github.com/mappu/miqt/qt"
)

type MenuView struct {
	*backbone.StatelessView
	model *Model
}

type MenuItem struct {
	Title     string
	OnClicked func()
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
		OnClicked: v.createNewGame,
	}))

	if v.model.SavesCount != 0 {
		layout.AddWidget(v.renderMenuItem(&MenuItem{
			Title:     "Continue",
			OnClicked: v.continueGame,
		}))

		layout.AddWidget(v.renderMenuItem(&MenuItem{
			Title: "Load Game",

			OnClicked: func() {
				router.Push(router.RouteSaves)
			},
		}))
	}

	layout.AddWidget(v.renderMenuItem(&MenuItem{
		Title:     "Exit",
		OnClicked: qt.QCoreApplication_Quit,
	}))

	return container
}

func (v *MenuView) renderMenuItem(item *MenuItem) *qt.QWidget {
	button := qt.NewQPushButton3(item.Title)
	button.OnClicked(item.OnClicked)
	button.SetStyleSheet(styled.Button)
	return button.QWidget
}

func (v *MenuView) createNewGame() {
	gameSave, err := v.model.NewGame()
	if err != nil {
		reporterr.Show(v.ViewRoot(), err)
	}

	router.Push(router.RouteGameWizard, router.Params{
		"gameSave": gameSave,
	})
}

func (v *MenuView) continueGame() {
	gameSave, err := v.model.GetLastGame()
	if err != nil {
		reporterr.Show(v.ViewRoot(), err)
		return
	}

	if gameSave.State != model.GameSaveReady {
		router.Push(router.RouteGameWizard, router.Params{
			"gameSave": gameSave,
		})
		return
	}

	router.Push(router.RouteGame, router.Params{
		"gameSave": gameSave,
	})
}
