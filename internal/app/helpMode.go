package app

import "charm.land/lipgloss/v2"

func HelpMode() error {
	drawHelp()
	return nil
}

func drawHelp() {
	help_output := []string{
		"taskjrnl - a simple command line task journal",
		"Usage:",
		"\ttaskjrnl|task [options] <command>",
		"\n",
		"Commands:",
		"\thelp\tShow help",
		"add",
		"\n",
		"Options:",
		"\t-v,--help\tShow help",
	}

	output := lipgloss.JoinVertical(lipgloss.Top, help_output...)
	println(output)
}
