package logs

import (
	"log/slog"
)

func AttrError(err error) slog.Attr {
	return slog.String("error", err.Error())
}
