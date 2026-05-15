package app

import (
	"database/sql"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
)

// Returns my ErrUsage error.
func NoCorrespondingMode(_ *sql.DB) error {
	return taskjrnlErrors.ErrUsage
}
