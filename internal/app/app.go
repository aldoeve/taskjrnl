package app

import (
	"flag"
	"fmt"
)

type AppOptions struct {
	help   bool
	list   bool
	add    bool
	done   bool
	jrnl   bool
	modify bool
}

func App() error {
	parameters := new(AppOptions)

	help := flag.Bool("help", false, "Show help")
	h := flag.Bool("h", false, "Show help")

	flag.Parse()

	positonalArgs := flag.Args()

	for index, arg := range positonalArgs {
		switch arg {
		case "help":
			parameters.help = true
		}
		fmt.Println(index)

	}

	parameters.help = *help || *h || parameters.help

	if parameters.help {
		DrawHelp()
	}

	return nil
}
