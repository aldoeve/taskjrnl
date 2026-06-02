package app

import (
	"database/sql"
	"strconv"
	store "taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"
)

func JrnlMode(db *sql.DB) error {
	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs != 2 {
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
