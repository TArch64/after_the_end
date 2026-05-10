package command

import (
	"slices"
	"sync"
)

var registry = &Registry{
	handlers: make(map[string][]*CmdHandler),
}

type Registry struct {
	handlers map[string][]*CmdHandler
	mx       sync.RWMutex
}

func (r *Registry) Read(kind string, reader func(handlers []*CmdHandler)) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	if handlers, ok := r.handlers[kind]; ok {
		reader(handlers)
	}
}

func (r *Registry) Add(kind string, handler *CmdHandler) {
	r.update(kind, func(handlers []*CmdHandler) []*CmdHandler {
		return append(handlers, handler)
	})
}

func (r *Registry) Delete(deleting *CmdHandler) {
	r.update(deleting.Kind, func(handlers []*CmdHandler) []*CmdHandler {
		return slices.DeleteFunc(handlers, func(handler *CmdHandler) bool {
			return handler.ID == deleting.ID
		})
	})
}

func (r *Registry) get(kind string) []*CmdHandler {
	if handlers, ok := r.handlers[kind]; ok {
		return handlers
	}
	return nil
}

func (r *Registry) update(kind string, update func(handlers []*CmdHandler) []*CmdHandler) {
	r.mx.Lock()
	defer r.mx.Unlock()
	r.handlers[kind] = update(r.get(kind))
}
