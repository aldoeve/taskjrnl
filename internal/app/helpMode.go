package app

import (
	"database/sql"
	appmodes "taskjrnl/internal/appModes"
	"taskjrnl/internal/config"
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
		argName := consts.HelpCommandsNFlagsTextStyle.Render("\t" + argObj.name)
		argUsageDesc := consts.HelpUsageTextStyle.Render(argObj.usage)
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
		{appmodes.Jrnl, "Add a note to a task. <taskNumber> [\"string\"]"},
		{appmodes.Modify, "Modify a task's values. <taskNumber> [priority:L|M|H] [tag:\"string\"]"},
		{appmodes.Weight, "Adjust the importance of a task. <taskNumber> <±numericValue>"},
	}

	return buildArgsNDesc(commandsSection, commands)
}

// Draws help output to stdout.
func drawHelp() {
	titleAppName := consts.HelpTitleStyle.Render("taskjrnl")
	titleAppDesc := consts.HelpOptionsTextStyle.Render(" — a simple command line task & journal")
	version := consts.HelpOptionsTextStyle.Render(config.Version)
	finalTitle := lipgloss.JoinHorizontal(lipgloss.Left, titleAppName, titleAppDesc, version)

	usageTabTitle := consts.HelpTitleStyle.Render("Usage:")
	usageText := consts.HelpUsageTextStyle.Render("\n\ttaskjrnl [options] <command>")
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

	final = consts.HelpBorderStyle.Render(final)

	lipgloss.Println(final)
}

// Calls the help drawer.
func HelpMode(_ *sql.DB) error {
	drawHelp()
	return nil
}
