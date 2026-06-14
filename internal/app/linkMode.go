package app

import (
	"database/sql"
	"strconv"
	"taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"
)

// Functions below this line are still not implemented.
func LinkMode(db *sql.DB) error {
	const (
		expectedNumArgs = 3
		taskALoc        = 0
		taskBLoc        = 1
		sharedNoteLoc   = 2
	)

	userInput := util.ArgsAfterKeyword()

	if numArgs := len(userInput); numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrUsage
	}

	userTaskAPosition, err := strconv.Atoi(userInput[taskALoc])
	if err != nil {
		return err
	}

	userTaskBPosition, err := strconv.Atoi(userInput[taskBLoc])
	if err != nil {
		return err
	}

	err = store.AreTasksInDB(db, userTaskAPosition, userTaskBPosition)
	if err == sql.ErrNoRows {
		return util.InformTasksDoesNotExist()
	}

	if err != nil {
		return err
	}

	err = store.AddNoteToTask(db, userTaskAPosition, userInput[sharedNoteLoc])
	if err != nil {
		return err
	}

	return store.AddNoteToTask(db, userTaskBPosition, userInput[sharedNoteLoc])
}
