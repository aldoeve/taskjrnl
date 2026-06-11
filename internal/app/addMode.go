// Pacage app contains the core application logic.
package app

import (
	"database/sql"
	store "taskjrnl/internal/store"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
	util "taskjrnl/pkg/util"
)

// Add logic to the application. Adds a task.
func AddMode(db *sql.DB) error {
	const (
		minNumArgs          = 1
		taskNameLoc         = minNumArgs - 1
		additionalParamsLoc = minNumArgs
	)

	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs < minNumArgs {
		return taskjrnlErrors.ErrTooFewArgs
	}

	if numArgs == minNumArgs {
		return store.CreateTask(db, userInput[taskNameLoc], nil, nil)
	}

	task, err := util.ParseTaskWithOptionalArgs(db, userInput[taskNameLoc], userInput[additionalParamsLoc:])
	if err != nil {
		return err
	}

	return store.CreateTask(db, task.Name, task.Tag, task.Priority)
}
