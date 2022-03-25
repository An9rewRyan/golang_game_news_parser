FROM golang:1.16-alpine

WORKDIR /go_parser

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY *.go ./


