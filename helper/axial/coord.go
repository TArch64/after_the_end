package axial

type Coord struct {
	Q int `json:"q" bun:",pk,notnull"`
	R int `json:"r" bun:",pk,notnull"`
	S int `json:"s" bun:",notnull"`
}
