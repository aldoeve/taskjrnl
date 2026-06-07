package app

import (
	"database/sql"
	store "taskjrnl/internal/store"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
	util "taskjrnl/pkg/util"
)

func ModifyMode(db *sql.DB) error {
	const (
		minNumArgs = 2
	)

	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs < minNumArgs {
		return taskjrnlErrors.ErrTooFewArgs
	}

	var err error

	if numArgs == 1 {
		err = store.CreateTask(db, userInput[0], nil, nil)
	} else {
		task, err := util.ParseTaskWithOptionalArgs(db, userInput[0], userInput[1:])
		if err != nil {
			return err
		}
		err = store.CreateTask(db, task.Name, task.Tag, task.Priority)
	}

	return err
}
