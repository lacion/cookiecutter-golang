# Change Log
All enhancements and patches to Cookiecutter Golang will be documented in this file.

## [2016-12-21]
### changed
- docker use now optional, you will get ask if you want to use docker while creating a project
- git use now optional as well
- added authors to credit contributors as well as general guidelines for contributing
- added testing
- viper config is now optional
- logrus is now optional

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
