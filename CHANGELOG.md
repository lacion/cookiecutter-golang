# Change Log
All enhancements and patches to Cookiecutter Golang will be documented in this file.

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
