# BUILDER
FROM golang:latest AS builder
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o abcd

# RUNNING
FROM debian:buster
RUN mkdir /app
WORKDIR /app/
CMD ["/app/abcd"]
