package backbone

import (
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
