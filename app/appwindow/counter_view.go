package appwindow

import (
	"fmt"

	"after_the_end/backbone"

	"github.com/mappu/miqt/qt"
)

type CounterView struct {
	*backbone.BaseView
	msg     *qt.QLabel
	counter int
}

func NewCounterView() *CounterView {
	return &CounterView{
		BaseView: backbone.NewBaseView(),
	}
}

func (v *CounterView) ViewInit(parent *qt.QWidget) {
	layout := qt.NewQVBoxLayout2()
	layout.AddStretch()

	v.msg = qt.NewQLabel3("Click the Button!")
	layout.AddWidget3(v.msg.QWidget, 0, qt.AlignCenter)

	btn := qt.NewQPushButton3("Click Me!")
	btn.SetFixedWidth(320)
	btn.OnPressed(v.onPressed)
	layout.AddWidget3(btn.QWidget, 0, qt.AlignCenter)

	layout.AddStretch()
	parent.SetLayout(layout.QLayout)
}

func (v *CounterView) ViewUpdate() {
	v.msg.SetText(fmt.Sprintf("You have clicked the button %d time(s)", v.counter))
}

func (v *CounterView) onPressed() {
	v.counter++
	v.ViewUpdate()
}
