# Builder
FROM golang:1.11.4-alpine3.8

RUN apk update && apk upgrade && \
    apk --update add git gcc make

WORKDIR /go/src/github.com/haffjjj/myblog-backend

COPY . .

EXPOSE 80

CMD make run