package cmdline

import (
	"flag"
	"fmt"
	"os"

	"github.com/rsmaxwell/job-to-xml/internal/basic"
)

// Config type
type CommandlineArguments struct {
	InputFilename  string
	OutputFilename string
}

func GetArguments() (CommandlineArguments, error) {

	help := flag.Bool("help", false, "Display help text")

	version := flag.Bool("version", false, "Display version")

	dsl := flag.String("dsl", "", "Input DSL filename")

	xml := flag.String("xml", "", "Output XML filename")

	flag.Parse()

	if *help {
		fmt.Printf("job-to-xml: Convert jenkins job dsl to the xml configuration for a SEED job\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *version {
		fmt.Printf("job-to-xml: Convert jenkins job dsl to the xml configuration for a SEED job\n")
		fmt.Printf("    Version:   %s\n", basic.Version())
		fmt.Printf("    BuildDate: %s\n", basic.BuildDate())
		fmt.Printf("    GitCommit: %s\n", basic.GitCommit())
		fmt.Printf("    GitBranch: %s\n", basic.GitBranch())
		fmt.Printf("    GitURL:    %s\n", basic.GitURL())
		os.Exit(0)
	}

	var args CommandlineArguments
	args.InputFilename = *dsl
	args.OutputFilename = *xml

	return args, nil
}
