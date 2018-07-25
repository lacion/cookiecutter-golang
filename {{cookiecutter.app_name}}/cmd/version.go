package cmd

import (
        "github.com/spf13/cobra"
        "fmt"
        "github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
        Use:   "version",
        Short: "Print the version number of {{cookiecutter.app_name}}",
        Long:  `All software has versions. This is {{cookiecutter.app_name}}`,
        Run: func(cmd *cobra.Command, args []string) {
                fmt.Println("Build Date:", version.BuildDate)
                fmt.Println("Git Commit:", version.GitCommit)
                fmt.Println("Version:", version.Version)
                fmt.Println("Go Version:", version.GoVersion)
                fmt.Println("OS / Arch:", version.OsArch)
        },
}

func init() {
        rootCmd.AddCommand(versionCmd)
}