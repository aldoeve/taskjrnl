// Pacage app contains the core application logic.
package app

import (
	"database/sql"
	"strings"
	"taskjrnl/internal/consts"
	errors "taskjrnl/internal/errors"
	store "taskjrnl/internal/store"
	util "taskjrnl/pkg/util"
)

func addTaskWithOptionalArgs(db *sql.DB, taskName string, optionalArgs []string) error {
	const priorityStr = "priority:"
	const tagStr = "tag:"

	var seenPriority bool
	var seenTag bool

	var priorityMode bool
	var tagMode bool

	var priorityValue string
	var tagValue string

	for _, arg := range optionalArgs {
		// Locates the keyword in a token and if its value is also with it.
		if !priorityMode && !tagMode {
			isPriority := strings.HasPrefix(arg, priorityStr)
			isTag := strings.HasPrefix(arg, tagStr)

			// Errors on no keywords or repeated keywords
			if (seenPriority && isPriority) || (seenTag && isTag) {
				return errors.IncorrectFormat
			}
			if !isTag && !isPriority {
				return errors.IncorrectFormat
			}

			value := strings.TrimPrefix(arg, priorityStr)
			value = strings.TrimPrefix(value, tagStr)

			seenPriority = seenPriority || isPriority
			seenTag = seenTag || isTag

			// Check to see if keyword's value is in the same token.
			// Otherwise, the next token's value wholy belongs to this keyword.
			// NOTE: Implies that tag can hold any value/keyword if malformed - intended that way.
			if len(value) > 0 {
				if isPriority {
					priorityValue = value
				} else {
					tagValue = value
				}
			} else {
				if isPriority {
					priorityMode = true
				} else {
					tagMode = true
				}
			}

			// Value to the keyword is in another token.
		} else {

			if priorityMode {
				value := strings.ToUpper(arg)
				if value == consts.LowPriority || value == consts.MidPriority || value == consts.HighPriority {
					priorityValue = value
					priorityMode = false
				} else {
					return errors.IncorrectFormat
				}
			}
			if tagMode {
				tagValue = arg
				tagMode = false
			}
		}
	}

	if priorityMode || tagMode {
		return errors.IncorrectFormat
	}

	// Sanitize
	dbReadyPriority := &priorityValue
	dbReadyTag := &tagValue

	if *dbReadyPriority == "" {
		dbReadyPriority = nil
	}
	if *dbReadyTag == "" {
		dbReadyTag = nil
	}

	return store.CreateTask(db, taskName, dbReadyTag, dbReadyPriority)
}

func AddMode(db *sql.DB) error {
	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs < 1 {
		return errors.ErrTooFewArgs
	}

	var err error

	if numArgs == 1 {
		err = store.CreateTask(db, userInput[0], nil, nil)
	} else {
		err = addTaskWithOptionalArgs(db, userInput[0], userInput[1:])
	}

	return err
}
