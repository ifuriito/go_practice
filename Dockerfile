FROM golang:1.16.3-alpine3.13
RUN apk update && apk add git
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
