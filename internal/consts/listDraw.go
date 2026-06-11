package consts

import (
	"taskjrnl/internal/config"
	initvars "taskjrnl/internal/initVars"

	"charm.land/lipgloss/v2"
)

const (
	ListCellLPad     = 0
	ListCellRPad     = 1
	ListWidth        = 14
	otherColumnWidth = 55
)

var (
	HeaderStyle    = lipgloss.NewStyle().Foreground(config.Vermilian).Bold(true).Align(lipgloss.Center)
	CellStyle      = lipgloss.NewStyle().Padding(ListCellLPad, ListCellRPad).Width(ListWidth)
	OddRowStyle    = lipgloss.NewStyle().Foreground(config.Orange)
	EvenRowStyle   = lipgloss.NewStyle().Foreground(config.Blue)
	TaskNameCutOff = initvars.DefaultTerminalWidth - otherColumnWidth
)
