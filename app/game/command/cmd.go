package command

import (
	"after_the_end/backbone"
	"after_the_end/helper/uniqid"
)

type handlerFn func(cmd Cmd)

var handlerId = uniqid.New()

type Cmd interface {
	Kind() string
}

type CmdHandler struct {
	ID       uniqid.ID
	Kind     string
	registry *Registry
	action   handlerFn
}

var _ backbone.Disposable = (*CmdHandler)(nil)

func (h *CmdHandler) Dispose() {
	h.registry.Delete(h)
}
