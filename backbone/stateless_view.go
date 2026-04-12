package backbone

import (
	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
)

type StatelessView struct {
	children []View
}

func NewStatelessView() *StatelessView {
	return &StatelessView{}
}

func (b *StatelessView) MountToWidget(parent *qt.QWidget, view View) {
	view.ViewBeforeInit()
	view.ViewInit(parent)
	b.children = append(b.children, view)
}

func (b *StatelessView) MountToLayout(layout *qt.QLayout, view View) {
	widget := qt.NewQWidget2()
	widget.SetObjectName("mount_layout")
	widget.SetStyleSheet(styled.Transparent)
	b.MountToWidget(widget, view)
	layout.AddWidget(widget)
}

func (b *StatelessView) ViewBeforeInit() {}

func (b *StatelessView) ViewUpdate() {}

func (b *StatelessView) ViewDestroy() {
	for _, child := range b.children {
		child.ViewDestroy()
	}

	if widget := b.Widget(); widget != nil {
		widget.DeleteLater()
	}

	if layout := b.Layout(); layout != nil {
		layout.DeleteLater()
	}
}

func (b *StatelessView) Widget() *qt.QWidget {
	return nil
}

func (b *StatelessView) Layout() *qt.QLayout {
	return nil
}
