# Start from golang v1.11 base image
FROM golang:1.11

# Add Maintainer Info
LABEL maintainer="Philip Porter"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/philaporter/random

# Copy everything
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Run the executable
CMD ["random"]