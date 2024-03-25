//go:generate go run gen.go

package buildinfo

import "fmt"

func PrintVersionInfo() {
	fmt.Printf("Version:   %s\n", Version())
	fmt.Printf("BuildDate: %s\n", BuildDate())
	fmt.Printf("GitCommit: %s\n", GitCommit())
	fmt.Printf("GitBranch: %s\n", GitBranch())
	fmt.Printf("GitURL:    %s\n", GitURL())
}
