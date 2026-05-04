// Package main contains the entry point of the application.
package main

import (
	"os"

	app "taskjrnl/internal/app"
	exitcodes "taskjrnl/internal/exitCodes"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
)

func checkForMinArgs() error {
	const MinCountArgs = 1
	numOfArgs := uint(len(os.Args[1:]))
	if numOfArgs < MinCountArgs {
		return taskjrnlErrors.ErrTooFewArgs
	}
	return nil
}

func main() {

	if err := checkForMinArgs(); err != nil {
		_ = app.HelpMode(nil)
		os.Exit(exitcodes.ExitUsage)
	}

	if err := app.App(); err != nil {
		_ = app.HelpMode(nil)
		os.Exit(exitcodes.ExitError)
	}

	os.Exit(exitcodes.ExitOk)
}
