package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const DOTNET_SDK = "/sdk"
const DOTNET_RUNTIME = "/shared/Microsoft.NETCore.App"
const ASPNET_RUNTIME = "/shared/Microsoft.AspNetCore.App"

func getDotnetBaseDir() string {
	basePath, err := exec.LookPath("dotnet")
	if err != nil {
		log.Fatalf("failed to find dotnet path: %s\n", err)
	}
	return basePath
}

func PrintDotnetDirs() {
	dotnetDir := getDotnetBaseDir()

	dotnetDir = strings.TrimSuffix(dotnetDir, "/dotnet")

	fmt.Printf("dotnet base directory: %s\n", dotnetDir)
	fmt.Printf("\tdotnet sdk directory: %s\n", dotnetDir+DOTNET_SDK)
	fmt.Printf("\tdotnet runtime directory: %s\n", dotnetDir+DOTNET_RUNTIME)
	fmt.Printf("\taspnet runtime directory: %s\n", dotnetDir+ASPNET_RUNTIME)
}
