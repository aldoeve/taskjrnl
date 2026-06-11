// Package main contains the entry point of the application.
package main

import (
	"os"

	app "taskjrnl/internal/app"
	exitcodes "taskjrnl/internal/exitCodes"
	initvars "taskjrnl/internal/initVars"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"

	"github.com/charmbracelet/x/term"
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
	initvars.TerminalWidth, initvars.TerminalHeight, initvars.TerminalError = term.GetSize(os.Stdout.Fd())
	if initvars.TerminalError != nil {
		initvars.TerminalHeight = initvars.DefaultTerminalHeight
		initvars.TerminalWidth = initvars.DefaultTerminalWidth
	}

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
