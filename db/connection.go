package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"

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
	params := url.Values{}
	params.Set("cache", "shared")
	params.Set("mode", "rwc")
	dataSource := fmt.Sprintf("file:%s?%s", path, params.Encode())
	sqlDB, err := sql.Open(sqliteshim.ShimName, dataSource)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(0)

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	if _, err = sqlDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}

	slog.Info("connected to database",
		slog.String("path", path),
	)

	return sqlDB, nil
}
