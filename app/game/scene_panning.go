package game

import (
	"github.com/mappu/miqt/qt"
)

type ScenePanning struct {
	lastPosition *qt.QPoint
	View         *qt.QGraphicsView
}

func NewScenePanning() *ScenePanning {
	return &ScenePanning{}
}

func (p *ScenePanning) isActive() bool {
	return p.lastPosition != nil
}

func (p *ScenePanning) Start(event *qt.QMouseEvent) {
	p.lastPosition = event.Pos()
	p.View.SetCursor(qt.NewQCursor2(qt.ClosedHandCursor))
}

func (p *ScenePanning) Move(event *qt.QMouseEvent) {
	if p.isActive() {
		current := event.Pos()
		delta := qt.NewQPoint3(current).OperatorMinusAssign(p.lastPosition)

		p.View.HorizontalScrollBar().SetValue(p.View.HorizontalScrollBar().Value() - delta.X())
		p.View.VerticalScrollBar().SetValue(p.View.VerticalScrollBar().Value() - delta.Y())

		p.lastPosition = current
	}
}

func (p *ScenePanning) End() {
	p.lastPosition = nil
	p.View.UnsetCursor()
}
