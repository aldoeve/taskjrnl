// Pacage app contains the core application logic.
package app

import (
	"database/sql"
	"slices"
	"strings"
	"taskjrnl/internal/consts"
	store "taskjrnl/internal/store"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
	util "taskjrnl/pkg/util"
)

const (
	priorityKeyword = "priority:"
	tagKeyword      = "tag:"
)

type mode int

const (
	noMode mode = iota
	priorityMode
	tagMode
)

type addParserState struct {
	mode mode

	seenPriority bool
	seenTag      bool

	finalPriorityVal string
	finalTagVal      string
}

// Figures out what keyword its looking at and sets the mode.
func (s *addParserState) figureOutMode(token string) {
	if strings.HasPrefix(token, priorityKeyword) {
		s.mode = priorityMode
		return
	}
	if strings.HasPrefix(token, tagKeyword) {
		s.mode = tagMode
	}
}

// Returns if any keyword has appeared more than once.
func (s *addParserState) hasRepeatedKeywords() bool {
	inPriority := (s.mode == priorityMode)
	inTag := (s.mode == tagMode)

	repeatedPriority := s.seenPriority && inPriority
	repeatedTag := s.seenTag && inTag

	s.seenPriority = s.seenPriority || inPriority
	s.seenTag = s.seenTag || inTag

	return repeatedPriority || repeatedTag
}

// Returns string with the seen keyword removed. I.E. The right hand-side of the keyword.
func (s *addParserState) stripKeyword(token string) string {
	if s.mode == priorityMode {
		return strings.TrimPrefix(token, priorityKeyword)
	}
	return strings.TrimPrefix(token, tagKeyword)
}

// Consumes a token and matches it to the corresponding mode.
func (s *addParserState) consumeAndAssign(value string) error {

	switch s.mode {
	case priorityMode:
		levelsOfPriority := []string{consts.LowPriority, consts.MidPriority, consts.HighPriority}
		if slices.Contains(levelsOfPriority, value) {
			s.finalPriorityVal = value
			s.mode = noMode
		} else {
			return taskjrnlErrors.IncorrectFormat
		}
	case tagMode:
		s.finalTagVal = value
		s.mode = noMode
	}

	return nil
}

// Basic Parser to figure out what to add to the database.
func addTaskWithOptionalArgs(db *sql.DB, taskName string, optionalArgs []string) error {

	state := addParserState{}

	for _, token := range optionalArgs {
		if state.mode == noMode {
			state.figureOutMode(token)

			if state.mode == noMode {
				return taskjrnlErrors.IncorrectFormat
			}
			if state.hasRepeatedKeywords() {
				return taskjrnlErrors.IncorrectFormat
			}

			value := state.stripKeyword(token)

			if value != "" {
				if state.consumeAndAssign(value) != nil {
					return taskjrnlErrors.IncorrectFormat
				}
			}

		} else {
			if state.consumeAndAssign(token) != nil {
				return taskjrnlErrors.IncorrectFormat
			}
		}
	}

	if state.mode != noMode {
		return taskjrnlErrors.IncorrectFormat
	}

	var (
		sanitizedPriority *string
		sanitizedTag      *string
	)

	if state.finalPriorityVal != "" {
		sanitizedPriority = &state.finalPriorityVal
	}
	if state.finalTagVal != "" {
		sanitizedTag = &state.finalTagVal
	}

	return store.CreateTask(db, taskName, sanitizedTag, sanitizedPriority)
}

// Add logic to the application. Adds a task.
func AddMode(db *sql.DB) error {
	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs < 1 {
		return taskjrnlErrors.ErrTooFewArgs
	}

	var err error

	if numArgs == 1 {
		err = store.CreateTask(db, userInput[0], nil, nil)
	} else {
		err = addTaskWithOptionalArgs(db, userInput[0], userInput[1:])
	}

	return err
}
