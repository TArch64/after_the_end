package cmd

import (
	"after_the_end/helper/axial"
)

type WalkPath struct {
	To *axial.Coord
}

func NewWalkPath(to *axial.Coord) *WalkPath {
	return &WalkPath{To: to}
}

func (cmd *WalkPath) Kind() string {
	return "WalkPath"
}
