# Go sandbox

Build and run the Docker image
$ docker build -t myapp .
$ docker run -it --rm myapp

Start an interactive shell and mount the project files so the source can be edited
outside the container.
$ docker run -v <full-path-to-project>:/go/src/app -i -t  myapp bash
