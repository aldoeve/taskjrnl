package app

import (
	"database/sql"
	"flag"
	appmodes "taskjrnl/internal/appModes"
	errors "taskjrnl/internal/errors"
	"taskjrnl/internal/store"
)

func bindStringToFunc(s *string) func(*sql.DB) error {
	modeHandlers := map[string]func(*sql.DB) error{
		appmodes.Add:    AddMode,
		appmodes.Done:   Done,
		appmodes.Help:   HelpMode,
		appmodes.Jrnl:   Jrnl,
		appmodes.List:   List,
		appmodes.Link:   Link,
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

	db, err := store.DBconnection()
	if err != nil {
		return err
	}
	defer db.Close()

	return mode(db)
}
func Done(_ *sql.DB) error {
	return nil
}
func Jrnl(_ *sql.DB) error {
	return nil
}
func List(_ *sql.DB) error {
	return nil
}
func Modify(_ *sql.DB) error {
	return nil
}
func Link(_ *sql.DB) error {
	return nil
}
