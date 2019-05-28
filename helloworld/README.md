## Reference
* https://grpc.io/docs/quickstart/go/
* https://medium.com/getamis/istio-基礎-grpc-負載均衡-d4be0d49ee07
* https://istio.io/docs/reference/config/networking/v1alpha3/virtual-service/
* https://medium.com/pismolabs/istio-envoy-grpc-metrics-winning-with-service-mesh-in-practice-d67a08acd8f7
* https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/grpc

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

TAG=v1.1
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

kubectl create namespace greeter
kubectl apply -f greeter-server.yaml
kubectl apply -f greeter-client.yaml

# (Optional) Change frontend service type to LoadBalancer
kubectl patch svc -n greeter greeter-client -p '{"spec":{"type":"LoadBalancer"}}'
```

## Run in Kubernetes, using Istio load balancer
```
git clone https://github.com/oneoneonepig/go-examples.git
cd go-examples/helloworld/kubernetes

kubectl create namespace greeter
kubectl label namespace greeter istio-injection=enabled

kubectl apply -f greeter-server.yaml
kubectl apply -f greeter-client.yaml

# Additional load balancing configuration, e.g. RANDOM selecting backend instead of default ROUND_ROBIN
kubectl apply -f istio-loadbalancer.yaml

# (Optional) Change frontend service type to LoadBalancer
kubectl patch svc -n greeter greeter-client -p '{"spec":{"type":"LoadBalancer"}}'
```

## Clean up
```
kubectl delete namespace greeter
```

