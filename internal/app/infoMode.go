package app

import (
	"database/sql"
	"strconv"
	"taskjrnl/internal/consts"
	"taskjrnl/internal/schema"
	"taskjrnl/internal/store"
	"taskjrnl/internal/taskjrnlErrors"
	"taskjrnl/pkg/util"

	"charm.land/lipgloss/v2"
)

// Formats and styles all the notes.
func formatTaskNotes(notes []schema.Pages) string {
	if len(notes) == 0 {
		return ""
	}

	var out string
	for index, value := range notes {
		var (
			date string
			text string
		)

		if index%2 == 0 {
			date = consts.InfoDateEvenRowStyle.Render(value.DateCreated)
			text = consts.InfoEvenRowStyle.Render(value.Note)
		} else {
			date = consts.InfoDateOddRowStyle.Render(value.DateCreated)
			text = consts.InfoOddRowStyle.Render(value.Note)
		}
		row := lipgloss.JoinHorizontal(lipgloss.Left, date, " -- ", text)
		out = lipgloss.JoinVertical(lipgloss.Left, out, row, "\n")
	}

	return out
}

// Formats and styles task info.
func formatTaskData(task schema.Tasks) string {
	date := consts.TaskDateStyle.Render(task.DateCreated)
	name := consts.TaskNameStyle.Render(task.Name)
	header := lipgloss.JoinHorizontal(lipgloss.Left, date, name)

	addtionalTaskInfo := consts.TaskDateStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left, " ",
		"Priority: ", *task.Priority, " ",
		"Weight: ", strconv.Itoa(task.ImportanceVariance),
	))

	return lipgloss.JoinVertical(lipgloss.Left, header, "\n", addtionalTaskInfo)
}

// Stiches all information together for output.
func renderInfo(task schema.Tasks, notes []schema.Pages) error {
	taskOutput := formatTaskData(task)
	notesOutput := formatTaskNotes(notes)

	out := taskOutput
	if notesOutput != "" {
		out = lipgloss.JoinVertical(lipgloss.Left, out, notesOutput)
	}

	out = consts.InfoBorder.Render(out)
	_, err := lipgloss.Println(out)

	return err
}

// Sanitizes the input and attempts to render desired task info.
func InfoMode(db *sql.DB) error {
	const (
		expectedNumArgs     = 1
		userTaskPositionLoc = 0
	)
	userInput := util.ArgsAfterKeyword()

	if numArgs := len(userInput); numArgs != expectedNumArgs {
		return taskjrnlErrors.ErrUsage
	}

	userTaskPosition, err := strconv.Atoi(userInput[userTaskPositionLoc])
	if err != nil {
		return err
	}

	task, err := store.FetchTaskinfo(db, userTaskPosition)
	if err == sql.ErrNoRows {
		return util.InformTasksDoesNotExist()
	} else if err != nil {
		return err
	}

	notes, err := store.FetchTaskNotes(db, task.Id)
	if err != nil {
		return err
	}

	return renderInfo(task, notes)
}
