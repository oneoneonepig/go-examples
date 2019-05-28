FROM golang:1.12.5

ADD . /go/src/github.com/oneoneonepig/go-examples/helloworld

RUN go get -u google.golang.org/grpc && \
    go install github.com/oneoneonepig/go-examples/helloworld/helloworld && \
    go install github.com/oneoneonepig/go-examples/helloworld/greeter_server && \
    go install github.com/oneoneonepig/go-examples/helloworld/greeter_client

CMD /go/bin/greeter_server
