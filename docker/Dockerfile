FROM golang:alpine

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /go/src/github.com/SND1231/post-service

RUN apk update \
 && apk add vim curl git make

COPY ./src .
COPY ./Makefile .
RUN git clone https://github.com/SND1231/proto-files.git \
 && mkdir proto \
 && cp -pr ./proto-files/post/* /go/src/github.com/SND1231/post-service/proto
RUN go get -u github.com/dgrijalva/jwt-go \
 && go get -u github.com/golang/protobuf \
 && go get -u github.com/gorilla/handlers \
 && go get -u github.com/grpc-ecosystem/grpc-gateway \
 && go get -u github.com/jinzhu/gorm \
 && go get -u golang.org/x/crypto \
 && go get -u golang.org/x/net \
 && go get -u google.golang.org/appengine \
 && go get -u google.golang.org/grpc
