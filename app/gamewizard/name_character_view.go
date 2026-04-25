package gamewizard

import (
	"errors"
	"strings"
	"time"

	"after_the_end/app/dialog/reporterr"
	"after_the_end/backbone"
	"after_the_end/backbone/styled"
	"after_the_end/helper/qttimer"

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
	onBack    func()
	onNext    func()
}

type NameCharacterViewOptions struct {
	Model  *GameSaveModel
	OnBack func()
	OnNext func()
}

func NewNameCharacterView(options *NameCharacterViewOptions) *NameCharacterView {
	return &NameCharacterView{
		StatelessView: backbone.NewStatelessView(),
		model:         options.Model,
		onBack:        options.OnBack,
		onNext:        options.OnNext,
	}
}

func (v *NameCharacterView) ViewInit() *qt.QWidget {
	widget := qt.NewQWidget2()
	widget.SetObjectName("name_character")

	column := qt.NewQVBoxLayout(widget)
	column.SetObjectName("name_character")
	column.SetSpacing(20)

	column.AddStretch()
	column.AddWidget(v.renderTitle())
	column.AddWidget(v.renderNameField())
	column.AddLayout(v.renderActions())
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
	v.nameField.SetPlaceholderText("Main Character Name")
	v.nameField.SetText(v.model.MainCharacterModel.Character.Name)

	go qttimer.In(100*time.Millisecond, func() {
		v.nameField.SetFocusWithReason(qt.OtherFocusReason)
	})

	return v.nameField.QWidget
}

func (v *NameCharacterView) renderActions() *qt.QLayout {
	row := qt.NewQHBoxLayout2()
	row.SetContentsMargins(0, 0, 0, 0)

	row.AddWidget2(v.renderAction(&NameCharacterAction{
		Title:     "Back",
		Name:      "name_character_back",
		OnClicked: v.onBack,
	}), 1)

	row.AddWidget2(v.renderAction(&NameCharacterAction{
		Title:     "Next",
		Name:      "name_character_next",
		OnClicked: v.nextStep,
	}), 3)

	return row.QLayout
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
		reporterr.Show(v.ViewRoot(), err)
		return
	}

	v.model.MainCharacterModel.Character.Name = v.nameField.Text()
	err = v.model.MainCharacterModel.Save("name")

	if err != nil {
		reporterr.Show(v.ViewRoot(), err)
		return
	}

	v.onNext()
}

func (v *NameCharacterView) validate() error {
	name := strings.TrimSpace(v.nameField.Text())
	v.nameField.SetText(name)

	if len(name) < 3 {
		return errors.New("character name cannot be shorter than 3 characters")
	}

	return nil
}
