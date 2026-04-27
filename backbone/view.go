package backbone

import (
	"after_the_end/helper/uniqid"

	qt "github.com/mappu/miqt/qt6"
)

type View interface {
	ViewID() uniqid.ID
	ViewRoot() *qt.QWidget
	ViewInit() *qt.QWidget
	ViewBeforeInit()
	ViewAfterInit(widget *qt.QWidget)
	ViewUpdate()
	ViewDestroy()
}
