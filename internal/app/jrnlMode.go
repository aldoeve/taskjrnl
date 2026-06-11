package app

import (
	"database/sql"
	"strconv"
	store "taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"
)

// App logic to add a note to a task.
func JrnlMode(db *sql.DB) error {
	const (
		expectedNumArgs     = 2
		userTaskPositionLoc = 0
	)

	userInput := util.ArgsAfterKeyword()

	if numArgs := len(userInput); numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrUsage
	}

	userTaskPosition, err := strconv.Atoi(userInput[userTaskPositionLoc])
	if err != nil {
		return err
	}

	err = store.AddNoteToTask(db, userTaskPosition, userInput[1])
	if err == sql.ErrNoRows {
		return util.InformTasksDoesNotExist()
	}

	return err
}
