package appwindow

import (
	"fmt"

	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type ButtonView struct {
	*backbone.BaseView
	btn     *qt.QPushButton
	counter int
}

func NewButtonView() *ButtonView {
	return &ButtonView{
		BaseView: backbone.NewBaseView(),
	}
}

func (v *ButtonView) ViewInit() error {
	v.btn = qt.NewQPushButton3("Hello world!")
	v.btn.SetFixedWidth(320)
	v.btn.OnPressed(v.onPressed)
	v.btn.Show()
	return nil
}

func (v *ButtonView) ViewUpdate() {
	v.btn.SetText(fmt.Sprintf("You have clicked the button %d time(s)", v.counter))
}

func (v *ButtonView) onPressed() {
	v.counter++
	v.ViewUpdate()
}
