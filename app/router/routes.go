package router

import (
	"after_the_end/backbone"
)

type RouteName string

const (
	RouteStart RouteName = "start"
	RouteSaves RouteName = "saves"
)

type Routes map[RouteName]backbone.View
