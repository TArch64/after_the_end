package reporterr

import (
	"log/slog"

	"github.com/mappu/miqt/qt"
)

func Show(parent *qt.QWidget, err error) {
	slog.Error("error report:",
		slog.String("msg", err.Error()),
	)

	qt.QMessageBox_Critical(parent, "Error!", err.Error())
}
