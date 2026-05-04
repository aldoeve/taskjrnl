package app

import (
	"database/sql"
	taskjrnlErrors "taskjrnl/internal/taskjrnlErrors"
)

func NoCorrespondingMode(_ *sql.DB) error {
	return taskjrnlErrors.ErrUsage
}
