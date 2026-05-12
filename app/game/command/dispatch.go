package command

func Dispatch[C Cmd](cmd C) {
	go registry.Read(cmd.Kind(), func(handlers []*CmdHandler) {
		for _, handler := range handlers {
			handler.action(cmd)
		}
	})
}

func BlockingDispatch[C WaitableCmd](cmd C) {
	Dispatch(cmd)
	cmd.Wait()
}
