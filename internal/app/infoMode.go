package app

import (
	"database/sql"
	"strconv"
	"taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"
)

func InfoMode(db *sql.DB) error {
	const expectedNumArgs = 1

	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrUsage
	}

	infoRequestedForTask, err := strconv.Atoi(userInput[0])
	if err != nil {
		return err
	}

	task, err := store.FetchTaskinfo(db, infoRequestedForTask)
	if err != nil {
		return err
	}

	_, err = store.FetchTaskNotes(db, task.Id)
	if err != nil {
		return err
	}

	return nil
}
