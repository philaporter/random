
# Start from golang v1.11 base image
FROM golang:1.11

# Add Maintainer Info
LABEL maintainer="Philip Porter"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/shutdown/app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 80

# Run the executable
CMD ["shutdown"]