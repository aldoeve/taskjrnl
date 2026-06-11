package consts

import (
	"taskjrnl/internal/config"

	"charm.land/lipgloss/v2"
)

const (
	CmdFlagWidth = 20
)

var (
	HelpTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(config.Vermilian)

	HelpUsageTextStyle = lipgloss.NewStyle().
				Foreground(config.Blue)

	HelpOptionsTextStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(config.Orange)

	HelpCommandsNFlagsTextStyle = HelpOptionsTextStyle.
					Width(CmdFlagWidth)

	HelpBorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(config.Vermilian)
)
