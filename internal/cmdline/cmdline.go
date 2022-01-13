package cmdline

import (
	"flag"
	"os"
)

// Config type
type CommandlineArguments struct {
	InputFilename  string
	OutputFilename string
}

func GetArguments() (CommandlineArguments, error) {

	help := flag.Bool("help", false, "Display help text")

	dsl := flag.String("dsl", "", "Input DSL filename")

	xml := flag.String("xml", "", "Output XML filename")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	var args CommandlineArguments
	args.InputFilename = *dsl
	args.OutputFilename = *xml

	return args, nil
}
