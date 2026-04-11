package router

var onPush func(name RouteName)

func Push(name RouteName) {
	onPush(name)
}
