package util

import (
	consts "taskjrnl/internal/consts"

	"charm.land/lipgloss/v2"
)

// Lets the users know no task exists.
func InformTasksDoesNotExist() error {
	noSuchTask := consts.DoneIssueTextStyle.Render("No such task found. Run [list] command to see tasks.")
	_, err := lipgloss.Println(noSuchTask)
	return err
}
