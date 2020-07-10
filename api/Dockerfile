FROM golang:1.12.5

COPY . /go

RUN go get github.com/nats-io/nats
RUN go build api.go

EXPOSE 8080
ENTRYPOINT ["/go/api"]