# set base image (host OS)
FROM golang:1.14
# set the working directory in the container
# Set the Current Working Directory inside the container

# Copy the binary produced by the docker instance
COPY src/ ./main




# command to run on container start
CMD [ "./main" ]
