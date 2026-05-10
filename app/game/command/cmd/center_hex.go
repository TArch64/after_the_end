package cmd

import (
	"after_the_end/app/game/command"
	"after_the_end/db/model"
)

type CenterHex struct {
	Coord *model.AxialCoord
}

var _ command.Cmd = (*CenterHex)(nil)

func (*CenterHex) Kind() string {
	return "CenterHex"
}
