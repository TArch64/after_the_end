package start

import (
	"context"
	"log/slog"

	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/db"
	"after_the_end/db/model"
	"after_the_end/logs"

	"github.com/mappu/miqt/qt"
)

type MenuView struct {
	*backbone.BaseView
}

type MenuItem struct {
	Title     string
	OnPressed func()
}

func NewMenuView() *MenuView {
	return &MenuView{
		BaseView: backbone.NewBaseView(),
	}
}

func (v *MenuView) ViewInit(parent *qt.QWidget) {
	layout := qt.NewQVBoxLayout2()
	layout.SetContentsMargins(0, 50, 0, 0)
	layout.SetObjectName("start_menu")
	parent.SetLayout(layout.QLayout)

	layout.AddWidget(v.renderMenuItem(&MenuItem{
		Title:     "New Game",
		OnPressed: v.newGame,
	}))

	layout.AddWidget(v.renderMenuItem(&MenuItem{
		Title: "Load Game",

		OnPressed: func() {
			router.Push(router.RouteSaves)
		},
	}))

	layout.AddWidget(v.renderMenuItem(&MenuItem{
		Title:     "Exit",
		OnPressed: qt.QCoreApplication_Quit,
	}))
}

func (v *MenuView) renderMenuItem(item *MenuItem) *qt.QWidget {
	button := qt.NewQPushButton3(item.Title)
	button.OnReleased(item.OnPressed)

	button.SetStyleSheet(`
    QPushButton {
        background-color: #dddddd;
        color: #444444;
        font-size: 18px;
        font-weight: bold;
        padding: 12px 32px;
        border: 3px solid #ffffff;
        border-right-color: #aaaaaa;
        border-bottom-color: #aaaaaa;
    }
    QPushButton:hover {
        background-color: #eeeeee;
        color: #000000;
    }
    QPushButton:pressed {
        background-color: #d5d5d5;
        border-top-color: #aaaaaa;
        border-left-color: #aaaaaa;
        border-right-color: #ffffff;
        border-bottom-color: #ffffff;
    }`)

	return button.QWidget
}

func (v *MenuView) newGame() {
	ctx := context.Background()

	savesCount, err := db.DB().
		NewSelect().
		Model((*model.GameSave)(nil)).
		Count(ctx)

	if err != nil {
		slog.Error("failed to create new game save",
			logs.AttrError(err),
		)
		return
	}

	save := &model.GameSave{
		Position: uint8(savesCount),
	}

	_, err = db.DB().
		NewInsert().
		Model(save).
		Exec(ctx)

	if err != nil {
		slog.Error("failed to create new game save",
			logs.AttrError(err),
		)
		return
	}

	slog.Info("successfully created new game save",
		slog.Uint64("save_id", uint64(save.ID)),
	)
}
