
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download 

COPY *.go ./cmd/ ./internal/

RUN go build -o yugioh-api ./cmd/api/

EXPOSE 8080

