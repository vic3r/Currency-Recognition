FROM golang:1.9.2

ENV GOBIN /go/bin

ADD . /go/src/github.com/vic3r/Currency-Recognition/
WORKDIR /go/src/github.com/vic3r/Currency-Recognition/

RUN go get -d -v github.com/gorilla/mux \
	&& go get -d -v gopkg.in/mgo.v2/bson \
	&& go get -d -v gopkg.in/mgo.v2

RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
dep ensure -v


WORKDIR /go/src/github.com/vic3r/Currency-Recognition/cmd/currency
RUN CGO_ENABLED=1 go build -i -o ../../Currency-Recognition
WORKDIR /go/src/github.com/vic3r/Currency-Recognition/
ENTRYPOINT ./Currency-Recognition
