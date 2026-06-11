package app

import (
	"database/sql"
	"strconv"
	store "taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"
)

// App logic to modify a tasks importance.
func WeightMode(db *sql.DB) error {
	const (
		expectedNumArgs     = 2
		userTaskPositionLoc = 0
		weightAdjustmentLoc = 1
	)

	userInput := util.ArgsAfterKeyword()

	if numArgs := len(userInput); numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrUsage
	}

	userTaskPosition, err := strconv.Atoi(userInput[userTaskPositionLoc])
	if err != nil {
		return taskjrnlErrors.ErrUsage
	}

	userWeightAdjustment, err := strconv.Atoi(userInput[weightAdjustmentLoc])
	if err != nil {
		return taskjrnlErrors.ErrUsage
	}

	err = store.AdjustTaskWeight(db, userTaskPosition, userWeightAdjustment)
	if err == sql.ErrNoRows {
		err = util.InformTasksDoesNotExist()
	}

	return err
}
