# cookiecutter-golang

[![Build Status](https://travis-ci.org/lacion/cookiecutter-golang.svg?branch=master)](https://travis-ci.org/lacion/cookiecutter-golang)

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter), Cookiecutter Golang is a framework for jumpstarting production-ready go projects quickly.

## Features

- Generous `Makefile` with management commands
- Uses [glide](https://github.com/Masterminds/glide) for dependency management

## Optional Integrations

- Can use [viper](https://github.com/spf13/viper) for env var config
- Can use [logrus](https://github.com/sirupsen/logrus) for logging
- Can creates dockerfile for building go binary and dockerfile for final go binary (no code in final container)
- If docker is used adds docker management commands to makefile

## Constraints

- Uses glide as only option for depency management
- Only maintained 3rd party libraries are used.

## Usage

Let's pretend you want to create a project called "echoserver". Rather than starting from scratch maybe copying 
some files and then editing the results to include your name, email, and various configuration issues that always 
get forgotten until the worst possible moment, get cookiecutter to do all the work.

First, get Cookiecutter. Trust me, it's awesome:
```console
$ pip install cookiecutter
```

alternatively you can install `cookiecutter` with homebrew:
```console
$ brew install cookiecutter
```

finally to run it based on this template just:
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
```

Enter the project and take a look around:
```console
$ cd echoserver/
$ ls
```

Run `make help` to see the available management commands, or just run `make` to build your project.
```console
$ make help
$ make
$ ./bin/echoserver
```

## Projects build with cookiecutter-golang

- [iothub](https://github.com/lacion/iothub) websocket multiroom server for IoT