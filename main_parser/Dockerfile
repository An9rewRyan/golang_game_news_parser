FROM golang:1.18
WORKDIR /go_parser

COPY go.mod ./
COPY go.sum ./

RUN apt-get update
RUN apt-get install -y apt-utils
RUN apt-get install -y firefox-esr
RUN apt-get install -y xvfb
RUN apt-get install -y default-jre
RUN apt-get install -y default-jdk

# RUN apt-get update
# RUN apt-get install -y firefox-esr


RUN go mod download
RUN go mod tidy
COPY *.go ./


