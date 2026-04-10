package backbone

import (
	"github.com/mappu/miqt/qt"
)

type View interface {
	ViewInit(parent *qt.QWidget)
	ViewUpdate()
	ViewDestroy()
}

type BaseView struct {
	children []View
}

func NewBaseView() *BaseView {
	return &BaseView{}
}

func (b *BaseView) Mount(parent *qt.QWidget, view View) {
	view.ViewInit(parent)
	b.children = append(b.children, view)
}

func (b *BaseView) ViewUpdate() {}

func (b *BaseView) ViewDestroy() {
	for _, child := range b.children {
		child.ViewDestroy()
	}
}
