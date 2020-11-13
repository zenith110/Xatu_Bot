# set base image (host OS)
FROM golang:1.14

# set the working directory in the container
WORKDIR /updater/
# Copy the binary produced by the docker instance
COPY src/ .

# command to run on container start
CMD [ "go", "run", "main.go"]
