package backbone

import (
	"github.com/mappu/miqt/qt"
)

type View interface {
	ViewInit(parent *qt.QWidget)
	ViewUpdate()
	ViewDestroy()
	Widget() *qt.QWidget
	Layout() *qt.QLayout
}

type BaseView struct {
	children []View
}

func NewBaseView() *BaseView {
	return &BaseView{}
}

func (b *BaseView) MountToWidget(parent *qt.QWidget, view View) {
	view.ViewInit(parent)
	b.children = append(b.children, view)
}

func (b *BaseView) MountToLayout(layout *qt.QLayout, view View) {
	widget := qt.NewQWidget2()
	widget.SetObjectName("mount_layout")
	widget.SetStyleSheet("background: transparent")
	b.MountToWidget(widget, view)
	layout.AddWidget(widget)
}

func (b *BaseView) ViewUpdate() {}

func (b *BaseView) ViewDestroy() {
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

func (b *BaseView) Widget() *qt.QWidget {
	return nil
}

func (b *BaseView) Layout() *qt.QLayout {
	return nil
}
