package consts

import (
	"taskjrnl/internal/config"

	"charm.land/lipgloss/v2"
)

var (
	HeaderStyle  = lipgloss.NewStyle().Foreground(config.Vermilian).Bold(true).Align(lipgloss.Center)
	CellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(14)
	OddRowStyle  = lipgloss.NewStyle().Foreground(config.Orange)
	EvenRowStyle = lipgloss.NewStyle().Foreground(config.Blue)
)
