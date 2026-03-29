// Pacage app contains the core application logic.
package app

import (
	"database/sql"
	"slices"
	"strings"
	"taskjrnl/internal/consts"
	errors "taskjrnl/internal/errors"
	store "taskjrnl/internal/store"
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

func (s *addParserState) figureOutMode(token string) {
	if strings.HasPrefix(token, priorityKeyword) {
		s.mode = priorityMode
		return
	}
	if strings.HasPrefix(token, tagKeyword) {
		s.mode = tagMode
	}
}

func (s *addParserState) hasRepeatedKeywords() bool {
	inPriority := (s.mode == priorityMode)
	inTag := (s.mode == tagMode)

	repeatedPriority := s.seenPriority && inPriority
	repeatedTag := s.seenTag && inTag

	s.seenPriority = s.seenPriority || inPriority
	s.seenTag = s.seenTag || inTag

	return repeatedPriority || repeatedTag
}

func (s *addParserState) stripKeyword(token string) string {
	if s.mode == priorityMode {
		return strings.TrimPrefix(token, priorityKeyword)
	}
	return strings.TrimPrefix(token, tagKeyword)
}

func (s *addParserState) consumeAndAssign(value string) error {

	switch s.mode {
	case priorityMode:
		levelsOfPriority := []string{consts.LowPriority, consts.MidPriority, consts.HighPriority}
		if slices.Contains(levelsOfPriority, value) {
			s.finalPriorityVal = value
			s.mode = noMode
		} else {
			return errors.IncorrectFormat
		}
	case tagMode:
		s.finalTagVal = value
		s.mode = noMode
	}

	return nil
}

func addTaskWithOptionalArgs(db *sql.DB, taskName string, optionalArgs []string) error {

	state := addParserState{}

	for _, token := range optionalArgs {
		if state.mode == noMode {
			state.figureOutMode(token)

			if state.mode == noMode {
				return errors.IncorrectFormat
			}
			if state.hasRepeatedKeywords() {
				return errors.IncorrectFormat
			}

			value := state.stripKeyword(token)

			if value != "" {
				if state.consumeAndAssign(value) != nil {
					return errors.IncorrectFormat
				}
			}

		} else {
			if state.consumeAndAssign(token) != nil {
				return errors.IncorrectFormat
			}
		}
	}

	if state.mode != noMode {
		return errors.IncorrectFormat
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
