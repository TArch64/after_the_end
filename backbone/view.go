package backbone

import (
	"context"

	"after_the_end/backbone/styled"

	"github.com/mappu/miqt/qt"
)

type View interface {
	ViewBeforeInit()
	ViewInit(parent *qt.QWidget)
	ViewUpdate()
	ViewDestroy()
	Widget() *qt.QWidget
	Layout() *qt.QLayout
}

type BaseView struct {
	children  []View
	Ctx       context.Context
	cancelCtx context.CancelFunc
}

func NewBaseView() *BaseView {
	return &BaseView{}
}

func (b *BaseView) ViewBeforeInit() {
	b.Ctx, b.cancelCtx = context.WithCancel(context.Background())
}

func (b *BaseView) MountToWidget(parent *qt.QWidget, view View) {
	view.ViewBeforeInit()
	view.ViewInit(parent)
	b.children = append(b.children, view)
}

func (b *BaseView) MountToLayout(layout *qt.QLayout, view View) {
	widget := qt.NewQWidget2()
	widget.SetObjectName("mount_layout")
	widget.SetStyleSheet(styled.Transparent)
	b.MountToWidget(widget, view)
	layout.AddWidget(widget)
}

func (b *BaseView) ViewUpdate() {}

func (b *BaseView) ViewDestroy() {
	b.cancelCtx()

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
