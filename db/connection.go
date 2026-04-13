package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"after_the_end/db/migrations"

	"github.com/mappu/miqt/qt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

var db *bun.DB

func Setup() error {
	databasePath, err := getDatabasePath()
	if err != nil {
		return err
	}

	sqlDB, err := openConnection(databasePath)
	if err != nil {
		return err
	}

	db = bun.
		NewDB(sqlDB, sqlitedialect.New()).
		WithQueryHook(bundebug.NewQueryHook(
			bundebug.WithVerbose(true),
		))

	return migrations.Up(db)
}

func DB() *bun.DB {
	return db
}

func getDatabasePath() (path string, err error) {
	path = qt.QStandardPaths_WritableLocation(qt.QStandardPaths__AppDataLocation)

	if err = os.MkdirAll(path, 0700); err != nil {
		path = ""
		return
	}

	path = filepath.Join(path, "application.db")
	return
}

func openConnection(path string) (*sql.DB, error) {
	dataSource := fmt.Sprintf("file:%s?cache=shared&mode=rwc", path)
	sqlDB, err := sql.Open(sqliteshim.ShimName, dataSource)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	slog.Info("connected to database",
		slog.String("path", path),
	)

	return sqlDB, nil
}
