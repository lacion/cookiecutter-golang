# cookiecutter-golang

.. image:: https://travis-ci.org/lacion/cookiecutter-golang.svg?branch=master
     :target: https://travis-ci.org/lacion/cookiecutter-golang?branch=master
     :alt: Build Status

A Go project template for [cookiecutter](https://github.com/audreyr/cookiecutter), Use it to make a new project built from scratch.

## Use it now

If you don't have `cookiecutter` installed, you'll need to install that first using Python's pip command:

```console
$ pip install cookiecutter
```

alternatively you can install `cookiecutter` via brew

```console
$ brew install cookiecutter
```

finally to run it based on this template just

```console
$ cookiecutter https://github.com/lacion/cookiecutter-golang.git
```

You will be asked about your basic info (name, project name, app name, etc.). This info will be used to customize your new project.

Run `make help` to see the available management commands, or just run `make` to build your project.

## Features

- Generous `Makefile` with management commands
- Works with either in a global `GOPATH` or with a local vendor directory
- Uses [glide](https://github.com/Masterminds/glide) for dependency management
- Uses [viper](https://github.com/spf13/viper) for env var config
- Uses [logrus](https://github.com/sirupsen/logrus) for logging

## Projects build with cookiecutter-golang

- [iothub](https://github.com/lacion/iothub) websocket multiroom server for IoT