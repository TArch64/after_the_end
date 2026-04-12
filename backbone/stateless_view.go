package backbone

import (
	"after_the_end/helper/uniqid"

	"github.com/mappu/miqt/qt"
)

var idFactory = uniqid.New()

type StatelessView struct {
	children map[uniqid.ID]View
	root     *qt.QWidget
	id       uniqid.ID
}

func NewStatelessView() *StatelessView {
	return &StatelessView{
		id: idFactory.Next(),
	}
}

func (b *StatelessView) ViewID() uniqid.ID {
	return b.id
}

func (b *StatelessView) Mount(view View) *qt.QWidget {
	if b.children == nil {
		b.children = make(map[uniqid.ID]View)
	}

	view.ViewBeforeInit()
	widget := view.ViewInit()
	b.children[view.ViewID()] = view
	view.ViewAfterInit(widget)
	return widget
}

func (b *StatelessView) Unmount(view View) {
	view.ViewDestroy()
	delete(b.children, view.ViewID())
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
