package command

import (
	"github.com/mappu/miqt/qt6/mainthread"
)

func Dispatch[C Cmd](cmd C) {
	go registry.Read(cmd.Kind(), func(handlers []*CmdHandler) {
		for _, handler := range handlers {
			handler.action(cmd)
		}
	})
}

func Handle[C Cmd](action func(cmd C)) *CmdHandler {
	handler := &CmdHandler{
		ID:       handlerId.Next(),
		registry: registry,

		action: func(cmd Cmd) {
			action(cmd.(C))
		},
	}

	var cmd C
	registry.Add(cmd.Kind(), handler)
	return handler
}

func MainThreadHandle[C Cmd](handle func(cmd C)) *CmdHandler {
	return Handle(func(cmd C) {
		mainthread.Start(func() {
			handle(cmd)
		})
	})
}
