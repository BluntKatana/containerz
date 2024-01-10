FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
RUN go build -v -o server

EXPOSE 1323

CMD ["./server"]