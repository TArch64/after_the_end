package confirm

import (
	"github.com/mappu/miqt/qt"
)

type Options struct {
	Parent *qt.QWidget
	Title  string
	Text   string
}

func Show(options *Options) bool {
	response := qt.QMessageBox_Question5(
		options.Parent,
		options.Title,
		options.Text,
		qt.QMessageBox__Yes|qt.QMessageBox__No,
	)

	return response == qt.QMessageBox__Yes
}
