# Build Stage
FROM {{cookiecutter.docker_build_image}}:{{cookiecutter.docker_build_image_version}} AS build-stage

LABEL app="build-{{cookiecutter.component_id}}"
LABEL REPO="https://github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}"

ENV PROJPATH=/go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}
WORKDIR /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}

RUN make build-alpine

# Final Stage
FROM {{cookiecutter.docker_image}}

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/{{cookiecutter.component_id}}/bin

WORKDIR /opt/{{cookiecutter.component_id}}/bin

COPY --from=build-stage /go/src/github.com/{{cookiecutter.github_username}}/{{cookiecutter.component_id}}/bin/{{cookiecutter.component_id}} /opt/{{cookiecutter.component_id}}/bin/
RUN chmod +x /opt/{{cookiecutter.component_id}}/bin/{{cookiecutter.component_id}}

# Create appuser
RUN adduser -D -g '' {{cookiecutter.component_id}}
USER {{cookiecutter.component_id}}

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/{{cookiecutter.component_id}}/bin/{{cookiecutter.component_id}}"]
