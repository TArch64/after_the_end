package qtgeometry

import (
	"github.com/mappu/miqt/qt"
)

func Read(target *qt.QWidget, onGeometry func(geometry *qt.QRect)) {
	onGeometry(target.Geometry())

	target.OnResizeEvent(func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent) {
		super(event)
		onGeometry(target.Geometry())
	})
}
