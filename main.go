package main

import (
	"os"

	"after_the_end/app/appwindow"
	"after_the_end/app/resources"
	_ "after_the_end/app/resources"
	"after_the_end/backbone/styled"
	"after_the_end/db"

	qt "github.com/mappu/miqt/qt6"
)

func main() {
	if err := db.Setup(); err != nil {
		panic(err)
	}

	app := qt.NewQApplication(os.Args)
	app.SetStyleSheet(styled.Global)
	initFont()

	windowView := appwindow.NewWindowView()
	windowView.ViewInit(nil)

	app.OnDestroyed(windowView.ViewDestroy)
	qt.QApplication_Exec()
}

func initFont() {
	fonts := []string{
		"Handjet-Black",
		"Handjet-Bold",
		"Handjet-ExtraBold",
		"Handjet-ExtraLight",
		"Handjet-Light",
		"Handjet-Medium",
		"Handjet-Regular",
		"Handjet-SemiBold",
		"Handjet-Thin",
	}

	for _, font := range fonts {
		qt.QFontDatabase_AddApplicationFont(resources.Font(font))
	}

	qt.QApplication_SetFont(qt.NewQFont6("Handjet", 20))
}
