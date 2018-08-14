FROM golang:latest
MAINTAINER Kris Nova "knova@heptio.com"
RUN mkdir -p /go/src/github.com/heptiolabs/git-events-operator
ADD . /go/src/github.com/heptiolabs/git-events-operator
WORKDIR /go/src/github.com/heptiolabs/git-events-operator
RUN go build -o git-events-operator .
CMD ["/go/src/github.com/heptiolabs/git-events-operator/git-events-operator"]