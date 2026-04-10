package main

import (
	"after_the_end/app/appwindow"
)

func main() {
	view := appwindow.NewWindowView()

	if err := view.ViewInit(); err != nil {
		panic(err)
	}
}
