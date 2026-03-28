// Pacage app contains the core application logic.
package app

import (
	"database/sql"
	errors "taskjrnl/internal/errors"
	store "taskjrnl/internal/store"
	util "taskjrnl/pkg/util"
)

func addTaskWithAddionalInfo(_ *sql.DB, _ []string) error {
	return nil
}

func AddMode(db *sql.DB) error {
	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs < 1 {
		return errors.ErrTooFewArgs
	}

	var err error

	if numArgs == 1 {
		err = store.CreateTask(db, userInput[0], nil, nil)
	} else {
		err = addTaskWithAddionalInfo(db, userInput)
	}

	return err
}
