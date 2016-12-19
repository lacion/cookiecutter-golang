FROM {{cookiecutter.docker_image}}

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/{{cookiecutter.app_name}}/bin

WORKDIR /opt/{{cookiecutter.app_name}}/bin

COPY bin/{{cookiecutter.app_name}} /opt/{{cookiecutter.app_name}}/bin/
RUN chmod +x /opt/{{cookiecutter.app_name}}/bin/{{cookiecutter.app_name}}

CMD /opt/{{cookiecutter.app_name}}/bin/{{cookiecutter.app_name}}