# cookiecutter-golang

[![Build Status](https://travis-ci.org/lacion/cookiecutter-golang.svg?branch=master)](https://travis-ci.org/lacion/cookiecutter-golang)

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter), Cookiecutter Golang is a framework for jumpstarting production-ready go projects quickly.

## Features

- Generous `Makefile` with management commands
- Uses `go dep` (with optional go module support *requires go 1.11*)
- injects build time and git hash at build time.

## Optional Integrations

- Can use [viper](https://github.com/spf13/viper) for env var config
- Can use [cobra](https://github.com/spf13/cobra) for cli tools
- Can use [logrus](https://github.com/sirupsen/logrus) for logging
- Can create dockerfile for building go binary and dockerfile for final go binary (no code in final container)
- If docker is used adds docker management commands to makefile
- Option of TravisCI, CircleCI or None

## Constraints

- Uses `dep` or `mod` for dependency management
- Only maintained 3rd party libraries are used.

This project now uses docker multistage builds, you need at least docker version v17.05.0-ce to use the docker file in this template, [you can read more about multistage builds here](https://www.critiqus.com/post/multi-stage-docker-builds/).

## Docker

This template uses docker multistage builds to make images slimmer and containers only the final project binary and assets with no source code whatsoever.

You can find the image dokcer file in this [repo](https://github.com/lacion/alpine-golang-buildimage) and more information about docker multistage builds in this [blog post](https://www.critiqus.com/post/multi-stage-docker-builds/).

Apps run under non root user and also with [dumb-init](https://github.com/Yelp/dumb-init).

## Usage

Let's pretend you want to create a project called "echoserver". Rather than starting from scratch maybe copying 
some files and then editing the results to include your name, email, and various configuration issues that always 
get forgotten until the worst possible moment, get cookiecutter to do all the work.

First, get Cookiecutter. Trust me, it's awesome:
```console
$ pip install cookiecutter
```

Alternatively, you can install `cookiecutter` with homebrew:
```console
$ brew install cookiecutter
```

Finally, to run it based on this template, type:
```console
$ cookiecutter https://github.com/lacion/cookiecutter-golang.git
```

You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.

Warning: After this point, change 'Luis Morales', 'lacion', etc to your own information.

Answer the prompts with your own desired [options](). For example:
```console
full_name [Luis Morales]: Luis Morales
github_username [lacion]: lacion
app_name [mygolangproject]: echoserver
project_short_description [A Golang project.]: Awesome Echo Server
docker_hub_username [lacion]: lacion
docker_image [lacion/docker-alpine:latest]: lacion/docker-alpine:latest
docker_build_image [lacion/docker-alpine:gobuildimage]: lacion/docker-alpine:gobuildimage
use_docker [y]: y
use_git [y]: y
use_logrus_logging [y]: y
use_viper_config [y]: y
use_cobra_cmd [y]: y
Select use_ci:
1 - travis
2 - circle
3 - none
Choose from 1, 2, 3 [1]: 1
```

Enter the project and take a look around:
```console
$ cd echoserver/
$ ls
```

Run `make help` to see the available management commands, or just run `make build` to build your project.
```console
$ make help
$ make build
$ ./bin/echoserver
```

## Projects build with cookiecutter-golang

- [iothub](https://github.com/lacion/iothub) websocket multiroom server for IoT
