
# build it:
# docker build -t backend-rest .
# run it:
# docker run --restart always -d backend-rest
# docker run -d backend-rest
# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang v1.12 base image
FROM golang:1.12

# Add Maintainer Info
LABEL maintainer="Hugo J. Bello <hjbello.wk@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $home/test

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8001

# Command to run the executable
CMD ["./main"]