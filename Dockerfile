# Using a compact OS
FROM golang:latest

MAINTAINER JJJabc <jx@wicwin.com>

# go get

RUN go get github.com/astaxie/beego
RUN go get github.com/garyburd/redigo/redis

# go run
RUN go run main.go

EXPOSE 8080
