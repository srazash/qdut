package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const DOTNET_SDK = "/sdk"
const DOTNET_RUNTIME = "/shared/Microsoft.NETCore.App"
const ASPNET_RUNTIME = "/shared/Microsoft.AspNetCore.App"

func getDotnetBaseDir() string {
	baseDir, err := exec.LookPath("dotnet")
	if err != nil {
		fmt.Printf("Unable to detect dotnet in the $PATH,")
		fmt.Printf("try passing in the path using the -p flag")
		log.Fatalf("%s\n", err)
	}
	return strings.TrimSuffix(baseDir, "/dotnet")
}

func isValidPDir(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}

func GetDotnetDirs() map[string]string {
	dotnetDirs := make(map[string]string, 3)

	dotnetBaseDir := getDotnetBaseDir()

	dotnetDirs["sdk"] = dotnetBaseDir + DOTNET_SDK
	dotnetDirs["dotnet_runtime"] = dotnetBaseDir + DOTNET_RUNTIME
	dotnetDirs["aspnet_runtime"] = dotnetBaseDir + ASPNET_RUNTIME

	return dotnetDirs
}

func PrintDotnetDirs() {
	dotnetDirs := GetDotnetDirs()

	fmt.Printf("dotnet base directory found in PATH: %s\n", getDotnetBaseDir())

	fmt.Printf("\tdotnet sdk directory: %s ? %t\n", dotnetDirs["sdk"], isValidPDir(dotnetDirs["sdk"]))
	fmt.Printf("\tdotnet runtime directory: %s ? %t\n", dotnetDirs["dotnet_runtime"], isValidPDir(dotnetDirs["dotnet_runtime"]))
	fmt.Printf("\taspnet runtime directory: %s ? %t\n", dotnetDirs["aspnet_runtime"], isValidPDir(dotnetDirs["aspnet_runtime"]))
}
