// Package util holds commonly used utility functions.
package util

import (
	"flag"
	"taskjrnl/internal/config"
	consts "taskjrnl/internal/consts"
	schema "taskjrnl/internal/schema"
	"time"
)

// Returns the flags with the first token discarded.
// This function assumes the first token is a keyword and discards it.
func ArgsAfterKeyword() []string {
	return flag.Args()[1:]
}

var priorityValue = map[string]int{
	consts.LowPriority:  0,
	consts.MidPriority:  1000,
	consts.HighPriority: 2000,
}

// Returns the importance of a task. The higher the number the more important.
func CalculateImportance(task *schema.Tasks) int {
	priority := priorityValue[*task.Priority]

	var daysSinceCreation int
	layout := config.TimeFormat // SQLite TEXT format.
	storedTime, err := time.Parse(layout, task.DateCreated)

	if err == nil {
		daysSinceCreation = int(time.Since(storedTime).Hours() / 24)
	}

	finalCalculation := priority + daysSinceCreation + task.ImportanceVariance

	return finalCalculation
}
