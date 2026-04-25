package qttimer

import (
	"time"

	"github.com/mappu/miqt/qt/mainthread"
)

func In(duration time.Duration, action func()) {
	timer := time.NewTimer(duration)
	<-timer.C
	mainthread.Start(action)
}
