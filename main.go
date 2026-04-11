package main

import (
	"os"

	"after_the_end/app/appwindow"
	_ "after_the_end/app/resources"
	"after_the_end/db"

	"github.com/mappu/miqt/qt"
)

func main() {
	if err := db.Setup(); err != nil {
		panic(err)
	}

	app := qt.NewQApplication(os.Args)
	initFont()

	windowView := appwindow.NewWindowView()
	windowView.ViewInit(nil)

	app.OnDestroyed(windowView.ViewDestroy)
	qt.QApplication_Exec()
}

func initFont() {
	qt.QFontDatabase_AddApplicationFont(":/fonts/pixelify-sans-bold.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/pixelify-sans-medium.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/pixelify-sans-regular.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/pixelify-sans-semibold.ttf")
	qt.QApplication_SetFont(qt.NewQFont6("Pixelify Sans", 20))
}
