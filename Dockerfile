FROM jguyomard/hugo-builder
ADD . /go/src/github.com/heptio/advocacy
WORKDIR /go/src/github.com/heptio/advocacy
CMD ["hugo"]


FROM golang:latest
MAINTAINER Kris Nova "knova@heptio.com"
RUN mkdir -p /go/src/github.com/heptiolabs/git-events-operator
ADD . /go/src/github.com/heptiolabs/git-events-operator
WORKDIR /go/src/github.com/heptiolabs/git-events-operator 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o git-events-operator .

FROM alpine:latest
RUN mkdir -p /go/src/github.com/heptiolabs/git-events-operator
WORKDIR /go/src/github.com/heptiolabs/git-events-operator
COPY --from=0 /go/src/github.com/heptiolabs/git-events-operator/git-events-operator .
CMD ["./git-events-operator"]