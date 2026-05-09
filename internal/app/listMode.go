package app

import (
	"database/sql"
	"strconv"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/schema"
	store "taskjrnl/internal/store"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
)

func drawTasks(tasks []schema.Tasks) error {
	headers := []string{"Position", "Priority", "Tag", "Task", "Date Created"}
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(consts.Vermilian)).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return consts.HeaderStyle
			case row%2 == 0:
				return consts.EvenRowStyle
			default:
				return consts.OddRowStyle
			}
		}).
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

	_, err := lipgloss.Println(t)
	return err
}

func ListMode(db *sql.DB) error {
	tasks, err := store.FetchAllTasks(db)
	if err != nil {
		return err
	}

	err = drawTasks(tasks)
	return err
}
