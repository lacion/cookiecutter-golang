package main

import (
	"flag"
	"fmt"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version"
)

func main() {

	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
        fmt.Println("Git Commit:", version.GitCommit)
        fmt.Println("Version:", version.Version)
        fmt.Println("Go Version:", version.GoVersion)
        fmt.Println("OS / Arch:", version.OsArch)
		return
	}

	fmt.Println("Hello.")
}
