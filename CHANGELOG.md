# Change Log
All enhancements and patches to Cookiecutter Golang will be documented in this file.

## [2019-12-09]
### changed
- updated golang docker base image version options
- made go mod the default option
- updated for gofmt -s format
- updated code to be go vet valid
- updated readme and contributors

## [2019-04-18]
### changed
- updated golang docker base image version options

## [2018-09-12]
### changed
- added option to use go mod instead of dep

## [2018-08-26]
### changed
- build image now defaults to go version 1.11

## [2018-08-09]
### changed
- new leaner base image

## [2018-08-09]
### changed
- golang 1.11 beta3 option

## [2018-08-08]
### changed
- run as not root user in docker
- run under dumb init

## [2018-07-31]
### changed
- better build image that now supports diferent golang versions [1.9.7, 1.10.3, 1.11beta2]

## [2018-07-25]
### changed
- fixed command to run test on the makefile
- refactor version into something a bit more useful

## [2018-05-21]
### changed
- upgraded circleci to 2.0
- fixed issues with docker build
- cobra cmd is now optional
- make file now injects build time and git hash correctly.

## [2018-05-21]
### changed
- moved from glide to dep

## [2017-06-06]
### changed
- updated CI templates to latest go versions
- updated docker file for new multistage builds

## [2016-12-21]
### changed
- docker use now optional, you will get ask if you want to use docker while creating a project
- git use now optional as well
- added authors to credit contributors as well as general guidelines for contributing
- added testing
- viper config is now optional
- logrus is now optional
- choice for CI tools (travis, circle or none)

## [2016-12-20]
### changed
- correct gitignore for vendor and bin
- updated readme

## [2016-12-19]
### changed
- post gen project hook to initialize git
- docker images configurable
- separated docker hub username from github username
- removed the use of os.exit to honor defers
- updated readme

## [2016-12-18]
### Start
- Initial release of Cookiecutter Golang project template
