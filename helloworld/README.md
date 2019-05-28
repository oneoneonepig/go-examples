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

## Build and push to Docker Hub
```
git clone https://github.com/oneoneonepig/go-examples.git
cd go-examples/helloworld

TAG=v1
docker build -t greeter:$TAG .
docker tag greeter:$TAG oneoneonepig/greeter:$TAG
docker push oneoneonepig/greeter:$TAG
```

## Run in local Docker container
```
git clone https://github.com/oneoneonepig/go-examples.git
cd go-examples/helloworld
docker build -t greeter:test .

docker run -d --name greeter-srv greeter:test
docker run -d -p 8080:8080 --name greeter-clt greeter:test /go/bin/greeter_client -host $(docker inspect --format='{{.NetworkSettings.IPAddress}}' greeter-srv) 

curl localhost:8080

```

## Run in Kubernetes
```
git clone https://github.com/oneoneonepig/go-examples.git
cd go-examples/helloworld/kubernetes

kubectl apply -f namespace.yaml
kubectl apply -f .

# (Optional) Change frontend service type to LoadBalancer
kubectl patch svc -n greeter greeter-client -p '{"spec":{"type":"LoadBalancer"}}'
```
