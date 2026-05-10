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
type Options = BaseArgStruct

// Returns a string of formated arguements and their description.
func buildArgsNDesc(finalOutput string, args []BaseArgStruct) string {
	for _, argObj := range args {
		argName := consts.HelpCommandsNFlagsText.Render("\t" + argObj.name)
		argUsageDesc := consts.HelpUsageText.Render(argObj.usage)
		argRow := lipgloss.JoinHorizontal(lipgloss.Left, argName, argUsageDesc)

		finalOutput = lipgloss.JoinVertical(lipgloss.Left, finalOutput, argRow)
	}

	return finalOutput
}

// Returns the help screen options section.
func buildOptionsSection() string {
	optionsSection := consts.HelpTitleStyle.Render("Options:")
	flags := []Options{
		{"-h, --help", "Show help"},
	}

	return buildArgsNDesc(optionsSection, flags)
}

// Returns the help screen commands section.
func buildCommnadsSection() string {
	commandsSection := consts.HelpTitleStyle.Render("Commands:")
	commands := []Command{
		{appmodes.Help, "Show help"},
		{appmodes.Add, "Adds task. <taskName> [priority:L|M|H] [tag:\"string\"]"},
		{appmodes.List, "Lists all tasks"},
	}

	return buildArgsNDesc(commandsSection, commands)
}

// Draws help output to stdout.
func drawHelp() {
	titleAppName := consts.HelpTitleStyle.Render("taskjrnl")
	titleAppDesc := consts.HelpOptionsText.Render(" — a simple command line task & journal")
	finalTitle := lipgloss.JoinHorizontal(lipgloss.Left, titleAppName, titleAppDesc)

	usageTabTitle := consts.HelpTitleStyle.Render("Usage:")
	usageText := consts.HelpUsageText.Render("\n\ttaskjrnl | task [options] <command>")
	finalUsage := lipgloss.JoinVertical(lipgloss.Left, usageTabTitle, usageText)

	commandsSection := buildCommnadsSection()

	optionsSection := buildOptionsSection()

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
