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
	const expectedNumArgs = 1

	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrTooFewArgs
	}

	userCompletedTaskId, err := strconv.Atoi(userInput[0])
	if err != nil {
		return taskjrnlErrors.ErrUsage
	}

	err = store.RemoveTask(db, userCompletedTaskId)
	if err == sql.ErrNoRows {
		err = util.InformTasksDoesNotExist()
	}
	if err != nil {
		return err
	}

	return nil
}
