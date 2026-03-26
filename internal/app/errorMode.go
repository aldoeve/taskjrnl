package app

import errors "taskjrnl/internal/errors"

func NoCorrespondingMode() error {
	return errors.ErrUsage
}
