FROM golang:1.20-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apk add --no-cache gcc musl-dev make
RUN go install github.com/cosmtrek/air@latest
