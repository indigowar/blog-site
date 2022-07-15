# syntax=docker/dockerfile:1

FROM golang:1.18 AS builder
WORKDIR /go/src/github.com/indigowar/blog-site/
COPY . .
RUN ["go", "mod", "download"]
RUN ["go", "build", "cmd/app/app.go"]
EXPOSE 8000
CMD ["./app"]
