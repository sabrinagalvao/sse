FROM golang:1.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]
