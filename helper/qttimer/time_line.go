package qttimer

import (
	"time"

	qt "github.com/mappu/miqt/qt6"
)

type Translation struct {
	Duration time.Duration
	Tick     func(step float64)
	Finish   func()
}

func TimeLine(translation *Translation) *qt.QTimeLine {
	timeLine := qt.NewQTimeLine2(int(translation.Duration.Milliseconds()))
	timeLine.SetEasingCurve(qt.NewQEasingCurve3(qt.QEasingCurve__InOutCubic))
	timeLine.GoGC()

	lastT := float64(0)
	timeLine.OnValueChanged(func(t float64) {
		translation.Tick(t - lastT)
		lastT = t
	})

	if translation.Finish != nil {
		timeLine.OnFinished(func() {
			translation.Finish()
		})
	}

	timeLine.Start()
	return timeLine
}
