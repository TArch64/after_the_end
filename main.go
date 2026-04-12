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
	qt.QFontDatabase_AddApplicationFont(":/fonts/Handjet-Black.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/Handjet-Bold.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/Handjet-ExtraBold.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/Handjet-ExtraLight.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/Handjet-Light.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/Handjet-Medium.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/Handjet-Regular.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/Handjet-SemiBold.ttf")
	qt.QFontDatabase_AddApplicationFont(":/fonts/Handjet-Thin.ttf")

	qt.QApplication_SetFont(qt.NewQFont6("Handjet", 20))
}
