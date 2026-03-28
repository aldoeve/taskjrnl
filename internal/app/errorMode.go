package app

import (
	"database/sql"
	errors "taskjrnl/internal/errors"
)

func NoCorrespondingMode(_ *sql.DB) error {
	return errors.ErrUsage
}
