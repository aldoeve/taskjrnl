package consts

import (
	"taskjrnl/internal/config"

	"charm.land/lipgloss/v2"
)

var (
	DoneIssueTextStyle = lipgloss.NewStyle().Foreground(config.Blue).
		Border(lipgloss.RoundedBorder()).BorderForeground(config.Vermilian)
)
