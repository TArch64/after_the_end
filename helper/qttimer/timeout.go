package qttimer

import (
	"time"

	qt "github.com/mappu/miqt/qt6"
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
