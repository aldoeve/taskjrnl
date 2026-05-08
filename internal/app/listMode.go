package app

import (
	"database/sql"
	"strconv"
	"taskjrnl/internal/schema"
	store "taskjrnl/internal/store"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
)

func drawTasks(tasks []schema.Tasks) {
	headers := []string{"Position", "Priority", "Tag", "Task", "Date Created"}
	t := table.New().
		Headers(headers...)

	// Tasks are already in order from most important to least.
	var position int
	for _, task := range tasks {
		position++

		var (
			priority string
			tag      string
		)

		if task.Priority != nil {
			priority = *task.Priority
		}
		if task.Tag != nil {
			tag = *task.Tag
		}

		t.Row(strconv.Itoa(position), priority, tag, task.Name, task.DateCreated)
	}
	lipgloss.Println(t)
}

func ListMode(db *sql.DB) error {
	tasks, err := store.FetchAllTasks(db)
	if err != nil {
		return err
	}

	drawTasks(tasks)
	return nil
}
