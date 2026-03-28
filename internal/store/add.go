package store

import "database/sql"

func AddSingleTask()

func AddSingleTaskNoOptions(_ *sql.DB, _ string) error {
	return nil
}
