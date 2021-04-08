FROM golang:1.16.3-alpine3.13
RUN apk update && apk add git
RUN apk add curl
CMD curl -D - -s  -o /dev/null http://example.com
RUN mkdir /go/src/app
WORKDIR /go/src/app
ADD . /go/src/app
