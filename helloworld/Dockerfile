FROM golang

ADD . /go/src/github.com/oneoneonepig/go-examples/helloworld

RUN go get -u google.golang.org/grpc
RUN go install github.com/oneoneonepig/go-examples/helloworld/helloworld
RUN go install github.com/oneoneonepig/go-examples/helloworld/greeter_server

ENTRYPOINT /go/bin/greeter_server
