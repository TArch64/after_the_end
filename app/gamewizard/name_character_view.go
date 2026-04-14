package gamewizard

import (
	"errors"
	"strings"

	"after_the_end/app/dialog/errorreport"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
)

type NameCharacterView struct {
	*backbone.StatelessView
	model     *Model
	nameField *qt.QLineEdit
}

func NewNameCharacterView(model *Model) *NameCharacterView {
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

	column.AddStretch()
	column.AddWidget3(v.renderTitle(), 0, qt.AlignCenter)
	column.AddWidget(v.renderNameField())
	column.AddWidget(v.renderNextButton())
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
	v.nameField.SetText(v.model.MainCharacter.Name)
	return v.nameField.QWidget
}

func (v *NameCharacterView) renderNextButton() *qt.QWidget {
	button := qt.NewQPushButton3("Next")
	button.SetObjectName("name_character_next")
	button.SetStyleSheet(styled.Button)

	button.OnClicked(v.nextStep)

	return button.QWidget
}

func (v *NameCharacterView) nextStep() {
	err := v.validate()
	if err != nil {
		errorreport.Show(v.ViewRoot(), err)
		return
	}

	err = v.model.UpdateMainCharacter(map[string]any{
		"name": v.nameField.Text(),
	})

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
