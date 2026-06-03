package app

import (
	"database/sql"
	"strconv"
	"taskjrnl/internal/config"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/schema"
	store "taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/table"
)

// Returns a pointer to a setup lipgloss table.
func generateTable() *table.Table {
	headers := []string{"Position", "Priority", "Tag", "Task", "Date Created"}
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(config.Vermilian)).
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

	return t
}

// Draws tasks to stdout.
func drawTasks(tasks []schema.Tasks) error {
	t := generateTable()

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

// Draws the tasks to the command line in order of most important first.
func ListMode(db *sql.DB) error {
	const expectedNumArgs = 0

	userInput := util.ArgsAfterKeyword()
	numArgs := len(userInput)

	if numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrUsage
	}

	tasks, err := store.FetchAllTasks(db)
	if err != nil {
		return err
	}

	err = drawTasks(tasks)
	return err
}
