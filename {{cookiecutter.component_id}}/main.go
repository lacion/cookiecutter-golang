package main

import (
	{% if cookiecutter.use_cobra_cmd == "n" %}"flag"
	"fmt"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}/version"{% endif %}
	{% if cookiecutter.use_cobra_cmd == "y" %}"github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}/cmd"{% endif %}
)

func main() {

    {% if cookiecutter.use_cobra_cmd == "y" %}
    cmd.Execute()
	{% else %}
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
    {% endif %}
}
