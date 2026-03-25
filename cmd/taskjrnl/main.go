package main

import (
	"os"
	exitcodes "taskjrnl/internal/exitCodes"
)

func main() {
	const minArgs uint = 1

	args := os.Args[1:]
	numArgs := uint(len(args))

	if numArgs < minArgs {
		drawHelp()
		os.Exit(exitcodes.ExitUsage)
	}

	var displayHelp bool
	var errorDetected bool

	for _, arg := range args {
		switch arg {
		case "-h", "--help", "help":
			displayHelp = true
		default:
			displayHelp = true
		}
	}

	if displayHelp {
		drawHelp()
	}

	exitCode := exitcodes.ExitOk
	if errorDetected {
		exitCode = exitcodes.ExitError
	}

	os.Exit(exitCode)
}
