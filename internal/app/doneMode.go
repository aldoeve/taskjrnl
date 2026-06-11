package app

import (
	"database/sql"
	"strconv"
	"taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"
)

// Deletion/completion logic of the application. Removes a task.
func DoneMode(db *sql.DB) error {
	const (
		expectedNumArgs        = 1
		userTaskPositionArgLoc = 0
	)

	userInput := util.ArgsAfterKeyword()

	if numArgs := len(userInput); numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrTooFewArgs
	}

	userTaskPosition, err := strconv.Atoi(userInput[userTaskPositionArgLoc])
	if err != nil {
		return taskjrnlErrors.ErrUsage
	}

	err = store.RemoveTask(db, userTaskPosition)

	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return util.InformTasksDoesNotExist()
	default:
		return err
	}
}
