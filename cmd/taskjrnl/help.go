package main

import "charm.land/lipgloss/v2"

func drawHelp() {
	help_output := []string{
		"taskjrnl - a simple command line task journal",
		"Usage:",
		"\ttaskjrnl|task <command> [options]",
		"\n",
		"Commands:",
		"\thelp\tShow help",
		"\n",
		"Options:",
		"\t-v,--help\tShow help",
	}

	output := lipgloss.JoinVertical(lipgloss.Top, help_output...)
	println(output)
}
