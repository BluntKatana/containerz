FROM golang:1.21

WORKDIR /usr/src/app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Build into a single binary
COPY db /usr/local/go/src/db
COPY handler /usr/local/go/src/handler
COPY server.go ./
RUN go build -v -o server

CMD ["./server"]