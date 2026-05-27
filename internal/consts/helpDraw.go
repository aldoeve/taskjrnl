package consts

import (
	"taskjrnl/internal/config"

	"charm.land/lipgloss/v2"
)

var (
	HelpTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(config.Vermilian)

	HelpUsageText = lipgloss.NewStyle().
			Foreground(config.Blue)

	HelpOptionsText = lipgloss.NewStyle().
			Bold(true).
			Foreground(config.Orange)

	HelpCommandsNFlagsText = HelpOptionsText.
				Width(20)

	HelpBorder = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(config.Vermilian)
)
