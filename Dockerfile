FROM golang:1.8.0-alpine
MAINTAINER Dongri Jin <dongrify@gmail.com>

RUN apk --no-cache add git

RUN go get github.com/kaneshin/lime

ADD . /go/src/github.com/dongri/rest-client
WORKDIR /go/src/github.com/dongri/rest-client/yourserver

EXPOSE 3001

CMD ["lime", "-bin=/tmp/yourserver", "-port=3001", "-app-port=3000"]
