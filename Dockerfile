# syntax=docker/dockerfile:1

FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /app/vanity-age

ENTRYPOINT [ "/app/vanity-age" ]

