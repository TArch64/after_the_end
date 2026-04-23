package mathg

func Abs[N Number](num N) N {
	if num < 0 {
		return -num
	}
	return num
}
