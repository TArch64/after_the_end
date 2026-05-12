package cmd

import (
	"after_the_end/app/game/command"
	"after_the_end/helper/axial"
)

type MoveMainCharacter struct {
	*command.WithCompletion
	Coord *axial.Coord
}

func NewMoveMainCharacter(coord *axial.Coord) *MoveMainCharacter {
	return &MoveMainCharacter{
		WithCompletion: command.NewWithCompletion(),
		Coord:          coord,
	}
}

var _ command.Cmd = (*MoveMainCharacter)(nil)

func (*MoveMainCharacter) Kind() string {
	return "MoveMainCharacter"
}
