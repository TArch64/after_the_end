package qttimer

import (
	"time"

	"github.com/mappu/miqt/qt"
)

func Timeout(duration time.Duration, action func()) {
	timer := qt.NewQTimer()
	timer.SetSingleShot(true)

	timer.OnTimeout(func() {
		action()
		timer.DeleteLater()
	})

	timer.Start(int(duration.Milliseconds()))
}
