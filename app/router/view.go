package router

import (
	"after_the_end/backbone"
)

type Route interface {
	backbone.View
	ViewBeforeOpen(params Params) error
}
