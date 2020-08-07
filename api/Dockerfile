FROM golang:latest

RUN mkdir /go/src/work

WORKDIR /go/src/work

ADD . /go/src/work
RUN go mod download

ENTRYPOINT ["go", "run", "main.go"]
