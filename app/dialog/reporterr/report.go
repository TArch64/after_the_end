package reporterr

import (
	"log/slog"

	qt "github.com/mappu/miqt/qt6"
)

func Show(parent *qt.QWidget, err error) {
	slog.Error("error report:",
		slog.String("msg", err.Error()),
	)

	qt.QMessageBox_Critical(parent, "Error!", err.Error())
}
