## Reference
https://grpc.io/docs/quickstart/go/

## Get gRPC
```
go get -u google.golang.org/grpc
```

## Get helloworld source
```
go get -d github.com/oneoneonepig/go-examples/helloworld
```

## Install protoc (Centos 7)
```
cd ~
PROTOC_ZIP=protoc-3.7.1-osx-x86_64.zip
curl -OL https://github.com/google/protobuf/releases/download/v3.7.1/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
rm -f $PROTOC_ZIP
```

## Install protoc (Debian 9)
```
sudo apt install golang-google-grpc-dev
```
## Generate gRPC code
```
cd ~/go/src/github.com/oneoneonepig/go-examples/helloworld
protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
```

## Start server
```
go run ./greeter_server/main.go
```

## Run client
```
go run ./greeter_client/main.go
```

## Containerize
```
```
