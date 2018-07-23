FROM golang:1.10.3

RUN mkdir -p /go/src/github.com/teguhmaulana25/shopee_tech
WORKDIR /go/src/github.com/teguhmaulana25/shopee_tech

ADD . /go/src/github.com/teguhmaulana25/shopee_tech


RUN go get -v