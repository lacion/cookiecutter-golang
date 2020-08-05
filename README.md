# cookiecutter-golang

[![Build Status](https://travis-ci.org/lacion/cookiecutter-golang.svg?branch=master)](https://travis-ci.org/lacion/cookiecutter-golang)

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter), Cookiecutter Golang is a framework for jumpstarting production-ready go projects quickly.

## Features

- Generous `Makefile` with management commands
- Uses `go dep` (with optional go module support _requires go 1.11_)
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

This template uses docker [multistage]](https://www.critiqus.com/post/multi-stage-docker-builds/) builds to make images slimmer and containers only the final project binary and assets with no source code whatsoever.

Apps run under non root user and also with [dumb-init](https://github.com/Yelp/dumb-init)

Run `make help` to see the available management commands, or just run `make build` to build your project.

```console
$ make help
$ make build
$ ./bin/echoserver
```

## Projects build with cookiecutter-golang

- [iothub](https://github.com/lacion/iothub) websocket multiroom server for IoT
