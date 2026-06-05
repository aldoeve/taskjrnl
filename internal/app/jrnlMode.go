package app

import (
	"database/sql"
	"strconv"
	store "taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"
)

func JrnlMode(db *sql.DB) error {
	const expectedNumArgs = 2

	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrUsage
	}

	taskToAddNoteID, err := strconv.Atoi(userInput[0])
	if err != nil {
		return err
	}
	err = store.AddNoteToTask(db, taskToAddNoteID, userInput[1])
	if err == sql.ErrNoRows {
		err = util.InformTasksDoesNotExist()
	}
	if err != nil {
		return err
	}
	return nil
}
