FROM golang:latest

WORKDIR /GoAhead

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build cmd/main.go

CMD ["./main"]