package app

import (
	"database/sql"
	"strconv"
	store "taskjrnl/internal/store"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
	util "taskjrnl/pkg/util"
)

// App logic to modify a task's values
func ModifyMode(db *sql.DB) error {
	const (
		minNumArgs            = 2
		userTaskPositionLoc   = 0
		userTaskArgsLocStart  = 1
		simpleNameEditRequest = 1
	)

	userInput := util.ArgsAfterKeyword()

	if numArgs := len(userInput); numArgs < minNumArgs {
		return taskjrnlErrors.ErrUsage
	}

	userTaskPosition, err := strconv.Atoi(userInput[userTaskPositionLoc])
	if err != nil {
		return taskjrnlErrors.ErrUsage
	}

	propertiesToModify := userInput[userTaskArgsLocStart:]

	if numArgs := len(propertiesToModify); numArgs == simpleNameEditRequest {
		err = store.ModifyTask(db, userTaskPosition, propertiesToModify[userTaskPositionLoc], nil, nil)
	} else {
		task, err := util.ParseTaskWithOptionalArgs(db, propertiesToModify[userTaskPositionLoc], propertiesToModify[userTaskArgsLocStart:])
		if err != nil {
			return err
		}

		err = store.ModifyTask(db, userTaskPosition, task.Name, task.Tag, task.Priority)
	}

	if err == sql.ErrNoRows {
		return util.InformTasksDoesNotExist()
	}

	return err
}
