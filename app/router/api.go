package router

var onPush func(name Name, params Params)

func Push(name Name, params ...Params) {
	if len(params) == 0 {
		onPush(name, nil)
	} else {
		onPush(name, params[0])
	}
}
