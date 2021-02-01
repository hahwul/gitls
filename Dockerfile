# BUILDER
FROM golang:latest AS builder
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o gitls

# RUNNING
FROM debian:buster
RUN mkdir /app
WORKDIR /app/
COPY --from=builder /go/src/app/gitls /app/gitls
CMD ["/app/gitls"]
