package consts

import (
	"taskjrnl/internal/config"

	"charm.land/lipgloss/v2"
)

const (
	InfoModeWidth       = 100
	AdditionalInfoWidth = 23
)

var (
	InfoBorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(config.Vermilian)

	TaskNameStyle        = lipgloss.NewStyle().Foreground(config.Blue).Bold(true).Align(lipgloss.Left).Width(InfoModeWidth)
	TaskDateStyle        = lipgloss.NewStyle().Foreground(config.Vermilian).Bold(true).Align(lipgloss.Center).Width(AdditionalInfoWidth)
	InfoDateOddRowStyle  = lipgloss.NewStyle().Foreground(config.Orange)
	InfoDateEvenRowStyle = lipgloss.NewStyle().Foreground(config.Blue)
	InfoOddRowStyle      = InfoDateOddRowStyle.Width(InfoModeWidth).Align(lipgloss.Left)
	InfoEvenRowStyle     = InfoDateEvenRowStyle.Width(InfoModeWidth).Align(lipgloss.Left)
)
