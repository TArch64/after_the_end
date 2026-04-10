package main

import (
	"os"

	"after_the_end/app/appwindow"

	"github.com/mappu/miqt/qt"
)

func main() {
	app := qt.NewQApplication(os.Args)

	windowView := appwindow.NewWindowView()
	windowView.ViewInit(nil)

	app.OnDestroyed(windowView.ViewDestroy)
	qt.QApplication_Exec()
}
