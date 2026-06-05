package consts

import (
	"taskjrnl/internal/config"

	"charm.land/lipgloss/v2"
)

var (
	InfoBorder = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(config.Vermilian)

	TaskName         = lipgloss.NewStyle().Foreground(config.Blue).Bold(true).Align(lipgloss.Center)
	TaskDate         = lipgloss.NewStyle().Foreground(config.Vermilian).Bold(true).Align(lipgloss.Center)
	InfoOddRowStyle  = lipgloss.NewStyle().Foreground(config.Orange)
	InfoEvenRowStyle = lipgloss.NewStyle().Foreground(config.Blue)
)
