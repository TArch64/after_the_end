package backbone

import (
	"github.com/mappu/miqt/qt"
)

type View interface {
	ViewInit() *qt.QWidget
	ViewBeforeInit()
	ViewAfterInit(widget *qt.QWidget)
	ViewUpdate()
	ViewDestroy()
}
