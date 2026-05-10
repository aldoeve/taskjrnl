package app

import (
	"database/sql"
	appmodes "taskjrnl/internal/appModes"
	"taskjrnl/internal/consts"

	"charm.land/lipgloss/v2"
)

type BaseArgStruct struct {
	name  string
	usage string
}
type Command = BaseArgStruct
type Flag = BaseArgStruct

// Draws help output to stdout.
func drawHelp() {
	titleAppName := consts.HelpTitleStyle.Render("taskjrnl")
	titleAppDesc := consts.HelpOptionsText.Render(" — a simple command line task & journal")
	finalTitle := lipgloss.JoinHorizontal(lipgloss.Left, titleAppName, titleAppDesc)

	usageTabTitle := consts.HelpTitleStyle.Render("Usage:")
	usageText := consts.HelpUsageText.Render("\n\ttaskjrnl | task [options] <command>")
	finalUsage := lipgloss.JoinVertical(lipgloss.Left, usageTabTitle, usageText)

	commandsSection := consts.HelpTitleStyle.Render("Commands:")
	commands := []Command{
		{appmodes.Help, "Show help"},
		{appmodes.Add, "Adds task. <taskName> [priority:L|M|H] [tag:\"string\"]"},
		{appmodes.List, "Lists all tasks"},
	}

	for _, commandObj := range commands {
		command := consts.HelpCommandsNFlagsText.Render("\t" + commandObj.name)
		commandUsageDesc := consts.HelpUsageText.Render(commandObj.usage)
		commandRow := lipgloss.JoinHorizontal(lipgloss.Left, command, commandUsageDesc)

		commandsSection = lipgloss.JoinVertical(lipgloss.Left, commandsSection, commandRow)
	}

	optionsSection := consts.HelpTitleStyle.Render("Options:")
	flags := []Flag{
		{"-h, --help", "Show help"},
	}

	for _, flagObj := range flags {
		flagName := consts.HelpCommandsNFlagsText.Render("\t" + flagObj.name)
		flagUsageDesc := consts.HelpUsageText.Render(flagObj.usage)
		flagRow := lipgloss.JoinHorizontal(lipgloss.Left, flagName, flagUsageDesc)

		optionsSection = lipgloss.JoinVertical(lipgloss.Left, optionsSection, flagRow)
	}

	final := lipgloss.JoinVertical(
		lipgloss.Left,
		finalTitle,
		"\n",
		finalUsage,
		"\n",
		commandsSection,
		"\n",
		optionsSection,
		"\n",
	)

	lipgloss.Println(final)
}

// Calls the help drawer.
func HelpMode(_ *sql.DB) error {
	drawHelp()
	return nil
}
