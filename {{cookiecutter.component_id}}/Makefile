{% if cookiecutter.use_docker == "y" %}
.PHONY: build build-alpine clean test help default
{% else %}
.PHONY: build clean test help default
{% endif %}


BIN_NAME={{cookiecutter.component_id}}

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')
IMAGE_NAME := "{{cookiecutter.docker_hub_username}}/{{cookiecutter.component_id}}"

default: test

help:
	@echo 'Management commands for {{cookiecutter.component_id}}:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	{% if cookiecutter.use_docker == "y" %}@echo '    make build-alpine    Compile optimized for alpine linux.'
	@echo '    make package         Build final docker image with just the go binary inside'
	@echo '    make tag             Tag image created by package with latest, git commit and version'
	@echo '    make test            Run tests on a compiled project.'
	@echo '    make push            Push tagged images to registry'{% endif %}
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

get-deps:
	dep ensure
{% if cookiecutter.use_docker == "y" %}
build-alpine:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}/version.BuildDate=${BUILD_DATE}' -o bin/${BIN_NAME}

package:
	@echo "building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):local .

tag: 
	@echo "Tagging: latest ${VERSION} $(GIT_COMMIT)"
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push: tag
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest
{% endif %}
clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...

