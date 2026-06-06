package consts

import (
	"taskjrnl/internal/config"

	"charm.land/lipgloss/v2"
)

const (
	InfoModeWidth = 100
)

var (
	InfoBorder = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(config.Vermilian)

	TaskNameStyle        = lipgloss.NewStyle().Foreground(config.Blue).Bold(true).Align(lipgloss.Center)
	TaskDateStyle        = lipgloss.NewStyle().Foreground(config.Vermilian).Bold(true).Align(lipgloss.Center)
	InfoDateOddRowStyle  = lipgloss.NewStyle().Foreground(config.Orange)
	InfoDateEvenRowStyle = lipgloss.NewStyle().Foreground(config.Blue)
	InfoOddRowStyle      = InfoDateOddRowStyle.Width(InfoModeWidth)
	InfoEvenRowStyle     = InfoDateEvenRowStyle.Width(InfoModeWidth)
)
