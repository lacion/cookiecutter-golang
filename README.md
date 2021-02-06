# cookiecutter-golang

[![Build Status](https://travis-ci.org/lacion/cookiecutter-golang.svg?branch=master)](https://travis-ci.org/lacion/cookiecutter-golang)

Powered by [Cookiecutter](https://github.com/audreyr/cookiecutter) - Cookiecutter Golang is a framework for jump-starting production-ready Go projects quickly.

## Features

- Generous `Makefile` with management commands
- Uses `go dep` (with optional [Go Modules](https://golang.org/ref/mod) support _requires go 1.11_)
- injects build time and git hash at build time.

## Optional Integrations

- Can use [viper](https://github.com/spf13/viper) for env var config
- Can use [cobra](https://github.com/spf13/cobra) for cli tools
- Can use [logrus](https://github.com/sirupsen/logrus) for logging
- Can create dockerfile for building go binary and dockerfile for final go binary (no code in final container)
- If Docker is used adds Docker management commands to makefile
- Option of TravisCI, CircleCI or None

## Constraints

- Uses `dep` or `mod` for dependency management.
- Only maintained 3rd party libraries are used.

This project now uses Docker multistage builds, you need at least Docker version v17.05.0-ce to use the Docker file in this template, [you can read more about multistage builds here](https://www.critiqus.com/post/multi-stage-docker-builds/).

## Docker

This template uses Docker [multi-stage](https://docs.docker.com/develop/develop-images/multistage-build) builds to make images slimmer and containers only the final project binary and assets with no source code whatsoever.

Apps run as non-root using [dumb-init](https://github.com/Yelp/dumb-init).

Run `make help` to see the available management commands, or just run `make build` to build your project.

```console
$ make help
$ make build
$ ./bin/echoserver
```

## Projects built with cookiecutter-golang

- [iothub](https://github.com/lacion/iothub) - multiroom websocket server for IoT
