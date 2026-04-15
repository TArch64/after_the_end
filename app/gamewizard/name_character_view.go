package gamewizard

import (
	"errors"
	"strings"

	"after_the_end/app/dialog/errorreport"
	"after_the_end/app/router"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
)

type NameCharacterAction struct {
	Title     string
	Name      string
	OnClicked func()
}

type NameCharacterView struct {
	*backbone.StatelessView
	model     *GameSaveModel
	nameField *qt.QLineEdit
}

func NewNameCharacterView(model *GameSaveModel) *NameCharacterView {
	return &NameCharacterView{
		StatelessView: backbone.NewStatelessView(),
		model:         model,
	}
}

func (v *NameCharacterView) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("name_character")

	column := qt.NewQVBoxLayout(widget)
	column.SetObjectName("name_character")
	column.SetSpacing(20)

	column.AddStretch()
	column.AddWidget3(v.renderTitle(), 0, qt.AlignCenter)
	column.AddWidget(v.renderNameField())
	column.AddWidget(v.renderActions())
	column.AddStretch()

	return widget
}

func (v *NameCharacterView) renderTitle() *qt.QWidget {
	title := qt.NewQLabel3("Name Your Stranger")
	title.SetObjectName("name_character_title")
	title.SetStyleSheet(styled.S("#name_character_title", styled.Title2))
	title.SetGraphicsEffect(styled.TitleShadow())
	return title.QWidget
}

func (v *NameCharacterView) renderNameField() *qt.QWidget {
	v.nameField = qt.NewQLineEdit2()
	v.nameField.SetObjectName("name_character_field")
	v.nameField.SetStyleSheet(styled.LineEdit)
	v.nameField.SetText(v.model.MainCharacterModel.Character.Name)
	return v.nameField.QWidget
}

func (v *NameCharacterView) renderActions() *qt.QWidget {
	widget := qt.NewQWidget2()
	row := qt.NewQHBoxLayout(widget)
	row.SetContentsMargins(0, 0, 0, 0)

	row.AddWidget2(v.renderAction(&NameCharacterAction{
		Title:     "Back",
		Name:      "name_character_back",
		OnClicked: v.backStep,
	}), 1)

	row.AddWidget2(v.renderAction(&NameCharacterAction{
		Title:     "Next",
		Name:      "name_character_next",
		OnClicked: v.nextStep,
	}), 3)

	return widget
}

func (v *NameCharacterView) renderAction(action *NameCharacterAction) *qt.QWidget {
	button := qt.NewQPushButton3(action.Title)
	button.SetObjectName(action.Name)
	button.SetStyleSheet(styled.Button)
	button.OnClicked(action.OnClicked)
	return button.QWidget
}

func (v *NameCharacterView) nextStep() {
	err := v.validate()
	if err != nil {
		errorreport.Show(v.ViewRoot(), err)
		return
	}

	v.model.MainCharacterModel.Character.Name = v.nameField.Text()
	err = v.model.MainCharacterModel.Save("name")

	if err != nil {
		errorreport.Show(v.ViewRoot(), err)
	}
}

func (v *NameCharacterView) validate() error {
	name := strings.TrimSpace(v.nameField.Text())
	v.nameField.SetText(name)

	if len(name) < 3 {
		return errors.New("character name cannot be shorter than 3 characters")
	}

	return nil
}

func (v *NameCharacterView) backStep() {
	router.Push(router.RouteStart)
}
