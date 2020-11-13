# set base image (host OS)
FROM golang:1.14
# set the working directory in the container
# Set the Current Working Directory inside the container

# copy the content of the local src directory to the working directory
COPY src/ .

# install dependencies
RUN go install .


# command to run on container start
CMD [ "go", "run", "./main.go" ]
