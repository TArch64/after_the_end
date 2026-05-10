package axial

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const CoordSeparator = ":"

var (
	InvalidCoordStr = errors.New("invalid coord string")
)

type Coord struct {
	Q int `json:"q" bun:",pk,notnull"`
	R int `json:"r" bun:",pk,notnull"`
}

func ParseCoord(str string) (*Coord, error) {
	parts := strings.SplitN(str, CoordSeparator, 2)
	if len(parts) != 2 {
		return nil, InvalidCoordStr
	}

	q, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, InvalidCoordStr
	}

	r, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, InvalidCoordStr
	}

	return &Coord{Q: q, R: r}, nil
}

func (c Coord) S() int {
	return -c.Q - c.R
}

func (c Coord) StringKey() string {
	return fmt.Sprintf("%d%s%d", c.Q, CoordSeparator, c.R)
}
