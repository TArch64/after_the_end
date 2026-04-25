package qttimer

func NextTick(action func()) {
	Timeout(0, action)
}
