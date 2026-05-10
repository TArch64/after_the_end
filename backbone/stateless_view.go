package backbone

import (
	"after_the_end/helper/uniqid"

	qt "github.com/mappu/miqt/qt6"
)

var viewID = uniqid.New()

type StatelessView struct {
	*DisposableView
	children map[uniqid.ID]View
	id       uniqid.ID
	root     *qt.QWidget
}

func NewStatelessView() *StatelessView {
	return &StatelessView{
		DisposableView: NewDisposableView(),
		id:             viewID.Next(),
	}
}

func (b *StatelessView) ViewID() uniqid.ID {
	return b.id
}

func (b *StatelessView) ViewRoot() *qt.QWidget {
	return b.root
}

func (b *StatelessView) Mount(view View, manualBeforeInit ...bool) *qt.QWidget {
	if b.children == nil {
		b.children = make(map[uniqid.ID]View)
	}

	if len(manualBeforeInit) == 0 || !manualBeforeInit[0] {
		view.ViewBeforeInit()
	}

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
