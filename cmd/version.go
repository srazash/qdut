package cmd

import "fmt"

const VERSION = "0.0.1"

func PrintAppVersion() {
	fmt.Printf("Quick (& Dirty) Dotnet Uninstall Tool v%s\n", VERSION)
}
