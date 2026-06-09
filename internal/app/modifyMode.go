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
		minNumArgs = 2
	)

	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs < minNumArgs {
		return taskjrnlErrors.ErrTooFewArgs
	}

	taskIdToModify, err := strconv.Atoi(userInput[0])
	if err != nil {
		return taskjrnlErrors.ErrUsage
	}

	propertiesToModify := userInput[1:]
	numArgs = len(propertiesToModify)

	if numArgs == 1 {
		err = store.ModifyTask(db, taskIdToModify, propertiesToModify[0], nil, nil)
	} else {
		task, err := util.ParseTaskWithOptionalArgs(db, propertiesToModify[0], propertiesToModify[1:])
		if err != nil {
			return err
		}
		err = store.ModifyTask(db, taskIdToModify, task.Name, task.Tag, task.Priority)
	}

	if err == sql.ErrNoRows {
		err = util.InformTasksDoesNotExist()
	}

	return err
}
