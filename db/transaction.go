package db

import (
	"errors"

	"github.com/uptrace/bun"
)

type TxFunc func(tx bun.Tx) error

func Tx(do TxFunc) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if err = do(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Join(err, rollbackErr)
		}
		return err
	}

	return tx.Commit()
}
