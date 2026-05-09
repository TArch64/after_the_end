package qtgeometry

import (
	qt "github.com/mappu/miqt/qt6"
)

type QResizable interface {
	Geometry() *qt.QRect
	OnResizeEvent(slot func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent))
}

func Read(target QResizable, onGeometry func(geometry *qt.QRect)) {
	onGeometry(target.Geometry())

	target.OnResizeEvent(func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent) {
		super(event)
		onGeometry(target.Geometry())
	})
}
