// Package main contains the entry point of the application.
package main

import (
	"os"

	app "taskjrnl/internal/app"
	exitcodes "taskjrnl/internal/exitCodes"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
)

// Returns error if there is no arguments passed.
func checkForMinArgs() error {
	const MinCountArgs = 1

	if numOfArgs := uint(len(os.Args[1:])); numOfArgs < MinCountArgs {
		return taskjrnlErrors.ErrTooFewArgs
	}

	return nil
}

// App entry point.
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
