package start

import (
	"after_the_end/backbone"

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
	layout.SetObjectName("start_menu")
	parent.SetLayout(layout.QLayout)

	layout.AddWidget(v.renderMenuItem(&MenuItem{
		Title:     "New Game",
		OnPressed: func() {},
	}))

	layout.AddWidget(v.renderMenuItem(&MenuItem{
		Title:     "Load Game",
		OnPressed: func() {},
	}))

	layout.AddWidget(v.renderMenuItem(&MenuItem{
		Title:     "Exit",
		OnPressed: qt.QCoreApplication_Quit,
	}))
}

func (v *MenuView) renderMenuItem(item *MenuItem) *qt.QWidget {
	button := qt.NewQPushButton3(item.Title)
	button.OnPressed(item.OnPressed)

	button.SetStyleSheet(`
		QPushButton {
				background-color: #E0E0E0;
				color: white;
				border: none;
				padding: 10px 20px;
		}
		QPushButton:hover {
				background-color: #B8B8B8;
		}
		QPushButton:pressed {
				background-color: #BFBFBF;
		}`)

	return button.QWidget
}
