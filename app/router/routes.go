package router

type Name string

const (
	RouteStart      Name = "start"
	RouteSaves      Name = "saves"
	RouteGameWizard Name = "game-wizard"
)

type Params map[string]any
type Routes map[Name]Route
