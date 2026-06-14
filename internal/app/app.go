package app

import (
	"database/sql"
	"flag"
	"path/filepath"
	appmodes "taskjrnl/internal/appModes"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/store"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"
)

// Returns the function that matches the string.
func bindStringToFunc(s *string) func(*sql.DB) error {
	strToAppModes := map[string]func(*sql.DB) error{
		appmodes.Add:    AddMode,
		appmodes.Done:   DoneMode,
		appmodes.Help:   HelpMode,
		appmodes.Jrnl:   JrnlMode,
		appmodes.Info:   InfoMode,
		appmodes.List:   ListMode,
		appmodes.Link:   LinkMode,
		appmodes.Modify: ModifyMode,
		appmodes.Weight: WeightMode,
	}

	if requestedFunc, ok := strToAppModes[*s]; ok {
		return requestedFunc
	}

	return NoCorrespondingMode
}

// Core appliction logic that swaps between modes.
func App() error {
	help := flag.Bool("help", false, "Show help")
	h := flag.Bool("h", false, "Show help")

	flag.Parse()

	var requestedMode string
	positonalArgs := flag.Args()

	if len(positonalArgs) > 0 {
		requestedMode = positonalArgs[0]
	}

	// Flags have mode-setting priority so they override values on purpose.

	if *help || *h {
		requestedMode = appmodes.Help
	}

	if requestedMode == "" {
		return taskjrnlErrors.ErrUsage
	}

	mode := bindStringToFunc(&requestedMode)

	appDir, err := util.CreateAppDir(consts.AppName)
	if err != nil {
		return err
	}

	dbLocation := filepath.Join(appDir, consts.DBName)

	db, err := store.DBconnection(dbLocation)
	if err != nil {
		return err
	}
	defer db.Close()

	return mode(db)
}
