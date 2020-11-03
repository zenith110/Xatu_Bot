# set base image (host OS)
FROM FROM golang:1.14
# set the working directory in the container
WORKDIR /code

# copy the content of the local src directory to the working directory
COPY src/ .

# install dependencies
RUN go get -d -v ./...
RUN go install -v ./...


# command to run on container start
CMD [ "go", "run", "./main.go" ]
