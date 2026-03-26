package app

import (
	"flag"
	appmodes "taskjrnl/internal/appModes"
	errors "taskjrnl/internal/errors"
)

func bindStringToFunc(s *string) func() error {
	modeHandlers := map[string]func() error{
		appmodes.Add:    Add,
		appmodes.Done:   Done,
		appmodes.Help:   HelpMode,
		appmodes.Jrnl:   Jrnl,
		appmodes.List:   List,
		appmodes.Modify: Modify,
	}
	if requestedFunc, found := modeHandlers[*s]; found {
		return requestedFunc
	}
	return NoCorrespondingMode
}

func App() error {
	help := flag.Bool("help", false, "Show help")
	h := flag.Bool("h", false, "Show help")

	flag.Parse()

	positonalArgs := flag.Args()
	numOfArgsLeft := len(positonalArgs)

	var requestedMode string

	if numOfArgsLeft > 0 {
		requestedMode = positonalArgs[0]
	}

	// Flags have mode-setting priority so they override values on purpose.

	if *help || *h {
		requestedMode = appmodes.Help
	}

	if requestedMode == "" {
		return errors.ErrUsage
	}

	mode := bindStringToFunc(&requestedMode)

	return mode()
}

func Add() error {
	return nil
}
func Done() error {
	return nil
}
func Jrnl() error {
	return nil
}
func List() error {
	return nil
}
func Modify() error {
	return nil
}
