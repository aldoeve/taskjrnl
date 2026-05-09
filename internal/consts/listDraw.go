package consts

import "charm.land/lipgloss/v2"

var (
	Vermilian = lipgloss.Color("#D55E00")
	Blue      = lipgloss.Color("#0072B2")
	Orange    = lipgloss.Color("#E69F00")

	HeaderStyle  = lipgloss.NewStyle().Foreground(Vermilian).Bold(true).Align(lipgloss.Center)
	CellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(14)
	OddRowStyle  = lipgloss.NewStyle().Foreground(Orange)
	EvenRowStyle = lipgloss.NewStyle().Foreground(Blue)
)
