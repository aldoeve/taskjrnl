package app

import (
	"database/sql"
	"fmt"
	"strconv"
	store "taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"
)

// App logic to modify a tasks importance.
func WeightMode(db *sql.DB) error {
	const expectedNumArgs = 2

	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrUsage
	}

	clientTaskPosition, err := strconv.Atoi(userInput[0])
	if err != nil {
		return taskjrnlErrors.ErrUsage
	}

	userValueAdjustment, err := strconv.Atoi(userInput[1])
	if err != nil {
		return taskjrnlErrors.ErrUsage
	}

	fmt.Println("read input")

	err = store.AdjustTaskWeight(db, clientTaskPosition, userValueAdjustment)
	if err == sql.ErrNoRows {
		err = util.InformTasksDoesNotExist()
	}
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
