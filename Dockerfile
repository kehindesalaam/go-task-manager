# golang image where workspace (GOPATH) configured at /go
FROM golang

# Copy the local package files to the container's workspace
ADD . /go/src/taskmanager

# Setting up working directory
WORKDIR /go/src/taskmanager

#Get godeps for managing and restoring dependencies
RUN go get github.com/tools/godep

#Restore godep dependencies
RUN godep restore

# Build the taskmanager command inside the container.
RUN go install taskmanager

# Run the taskmanager command when the container starts.
ENTRYPOINT /go/bin/taskmanager

# Service listens on port 3000.
EXPOSE 3000