package scene

import (
	"github.com/mappu/miqt/qt"
)

type Panning struct {
	lastPosition *qt.QPoint
	View         *qt.QGraphicsView
}

func NewPanning() *Panning {
	return &Panning{}
}

func (p *Panning) isActive() bool {
	return p.lastPosition != nil
}

func (p *Panning) Start(event *qt.QMouseEvent) {
	p.lastPosition = event.Pos()
	p.View.SetCursor(qt.NewQCursor2(qt.ClosedHandCursor))
}

func (p *Panning) Move(event *qt.QMouseEvent) {
	if p.isActive() {
		current := event.Pos()
		delta := qt.NewQPoint3(current).OperatorMinusAssign(p.lastPosition)

		p.View.HorizontalScrollBar().SetValue(p.View.HorizontalScrollBar().Value() - delta.X())
		p.View.VerticalScrollBar().SetValue(p.View.VerticalScrollBar().Value() - delta.Y())

		p.lastPosition = current
	}
}

func (p *Panning) End() {
	p.lastPosition = nil
	p.View.UnsetCursor()
}
