package backbone

import (
	"github.com/mappu/miqt/qt"
)

type StatelessView struct {
	children []View
	root     *qt.QWidget
}

func NewStatelessView() *StatelessView {
	return &StatelessView{}
}

func (b *StatelessView) Mount(view View) *qt.QWidget {
	view.ViewBeforeInit()
	widget := view.ViewInit()
	b.children = append(b.children, view)
	view.ViewAfterInit(widget)
	return widget
}

func (b *StatelessView) ViewBeforeInit() {}

func (b *StatelessView) ViewUpdate() {}

func (b *StatelessView) ViewAfterInit(widget *qt.QWidget) {
	b.root = widget
}

func (b *StatelessView) ViewDestroy() {
	for _, child := range b.children {
		child.ViewDestroy()
	}

	b.children = nil
	b.root.DeleteLater()
	b.root = nil
}
