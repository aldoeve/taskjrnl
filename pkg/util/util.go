// Package util holds commonly used utility functions.
package util

import (
	"flag"
	consts "taskjrnl/internal/consts"
	schema "taskjrnl/internal/schema"
	"time"
)

// This function assumes the first token is a keyword and discards it.
func ArgsAfterKeyword() []string {
	return flag.Args()[1:]
}

var PriorityValue = map[string]int{
	consts.LowPriority:  0,
	consts.MidPriority:  1000,
	consts.HighPriority: 2000,
}

func CalculateImportance(task *schema.Tasks) int {
	priority := PriorityValue[*task.Priority]

	var daysSinceCreation int
	layout := "2009-09-25 15:09:01" // SQLite TEXT format.
	storedTime, err := time.Parse(layout, task.DateCreated)
	if err == nil {
		daysSinceCreation = int(time.Since(storedTime).Hours() / 24)
	}

	finalCalculation := priority + daysSinceCreation + task.ImportanceVariance

	return finalCalculation
}
