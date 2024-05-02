package main

import (
	"fmt"

	"github.com/aatumaykin/go-simple-version"
)

func main() {
	info := version.Get()
	fmt.Println()
	fmt.Println("Version: ", info.Version)
	fmt.Println("Git Commit: ", info.Commit)
	fmt.Println("Build Date: ", info.BuildTime)
	fmt.Println("Go Version: ", info.GoVersion)
	fmt.Println("OS: ", info.Os)
	fmt.Println("Arch: ", info.Arch)
}
