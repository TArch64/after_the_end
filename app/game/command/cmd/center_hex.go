package cmd

import (
	"after_the_end/app/game/command"
	"after_the_end/db/model"
)

type ActivateHex struct {
	Coord *model.AxialCoord
}

var _ command.Cmd = (*ActivateHex)(nil)

func (*ActivateHex) Kind() string {
	return "CenterHex"
}
