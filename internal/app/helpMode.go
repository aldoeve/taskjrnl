package app

import (
	"database/sql"
	"fmt"
	appmodes "taskjrnl/internal/appModes"

	"charm.land/lipgloss/v2"
)

type BaseArgStruct struct {
	name  string
	usage string
}
type Command = BaseArgStruct
type Flag = BaseArgStruct

func HelpMode(_ *sql.DB) error {
	drawHelp()
	return nil
}

func drawHelp() {
	const defaultPadding = 10

	commands := []Command{
		{appmodes.Help, "Show help"},
		{appmodes.Add, "Adds task. <taskName> [priority:L|M|H] [tags:\"string\"] "},
	}
	flags := []Flag{
		{"-h,--help", "Show Help"},
	}

	helpOutput := []string{
		"\ntaskjrnl - a simple command line task journal",
		"Usage:",
		"\ttaskjrnl|task [options] <command>",
		"\n",
		"Commands:",
	}

	formatCMDorFlagArrays := func(array []BaseArgStruct, padding int) []string {
		var paddedStrings []string
		for _, arg := range array {
			paddedStrings = append(paddedStrings, fmt.Sprintf("\t%-*s %-s", padding, arg.name, arg.usage))
		}
		return paddedStrings
	}

	helpOutput = append(helpOutput, formatCMDorFlagArrays(commands, defaultPadding)...)
	helpOutput = append(helpOutput, "Options:")
	helpOutput = append(helpOutput, formatCMDorFlagArrays(flags, defaultPadding)...)
	helpOutput = append(helpOutput, "\n")

	output := lipgloss.JoinVertical(lipgloss.Top, helpOutput...)
	println(output)
}
