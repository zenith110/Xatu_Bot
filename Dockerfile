# set base image (host OS)
FROM golang:1.14-alpine3.12

# set the working directory in the container
WORKDIR /src/

# Copy the binary produced by the docker instance
RUN go get github.com/bwmarrin/discordgo
COPY src/ .


# command to run on container start
CMD [ "go", "run", "main.go"]
