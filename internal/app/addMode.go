// Pacage app contains the core application logic.
package app

import (
	"database/sql"
	errors "taskjrnl/internal/errors"
	util "taskjrnl/pkg/util"
)

func AddMode(db *sql.DB) error {
	userInput := util.ArgsAfterKeyword()

	if len(userInput) < 1 {
		return errors.ErrTooFewArgs
	}

	return nil
}
