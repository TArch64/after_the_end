package scene

import (
	qt "github.com/mappu/miqt/qt6"
)

type Panning struct {
	lastPosition *qt.QPoint
	isActive     bool
	View         *qt.QGraphicsView
}

func NewPanning() *Panning {
	return &Panning{}
}

func (p *Panning) Start(event *qt.QMouseEvent) {
	p.lastPosition = event.Pos()
}

func (p *Panning) Move(event *qt.QMouseEvent) {
	if p.lastPosition == nil {
		return
	}

	current := event.Pos()
	delta := qt.NewQPoint3(current).OperatorMinusAssign(p.lastPosition)

	if !p.isActive {
		p.isActive = delta.ManhattanLength() > 5

		if p.isActive {
			p.View.SetCursor(qt.NewQCursor2(qt.ClosedHandCursor))
		}
	}

	if p.isActive {
		p.View.HorizontalScrollBar().SetValue(p.View.HorizontalScrollBar().Value() - delta.X())
		p.View.VerticalScrollBar().SetValue(p.View.VerticalScrollBar().Value() - delta.Y())
		p.lastPosition = current
	}
}

func (p *Panning) End() bool {
	p.lastPosition = nil
	wasActive := p.isActive

	if wasActive {
		p.View.UnsetCursor()
		p.isActive = false
	}

	return wasActive
}
