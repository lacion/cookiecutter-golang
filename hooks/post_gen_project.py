"""
Does the following:

1. Inits git if used
2. Deletes dockerfiles if not going to be used
"""
import os
from subprocess import Popen

# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

def init_git():
    """
    Initialises git on the new project folder
    """
    GIT_COMMANDS = [
        ["git", "init"],
        ["git", "add", "."],
        ["git", "commit", "-a", "-m", "Initial Commit."]
    ]
    for command in GIT_COMMANDS:
        git = Popen(command, cwd=PROJECT_DIRECTORY)
        git.wait()


def remove_docker_files():
    """
    Removes files needed for docker if it isn't going to be used
    """
    for filename in ["Dockerfile", "Dockerfile.build"]:
        os.remove(os.path.join(
            PROJECT_DIRECTORY, filename
        ))

# 1. Initialize Git
if '{{ cookiecutter.use_docker }}'.lower() != 'y':
    remove_docker_files()

# 1. Remove Dockerfiles if docker is not going to be used
if '{{ cookiecutter.use_docker }}'.lower() != 'y':
    remove_docker_files()
