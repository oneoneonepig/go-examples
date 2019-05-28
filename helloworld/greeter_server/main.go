/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
        "context"
        "flag"
        "log"
        "net"
        "os"

        "google.golang.org/grpc"
        pb "github.com/oneoneonepig/go-examples/helloworld/helloworld"
)

var (
        host string
        port string
	nodeName string
	podName string
)

// Initialize listen port and address
func init() {
        flag.StringVar(&host, "host", "0.0.0.0", "Listening address")
        flag.StringVar(&port, "port", "3000", "Listening port")
        flag.Parse()
        nodeName = os.Getenv("NODE_NAME")
        podName = os.Getenv("POD_NAME")
	if len(nodeName) == 0 {
		nodeName = "NULL"
	}
	if len(podName) == 0 {
		podName = "NULL"
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
        log.Printf("Received: %v, from %v", in.Name, in.Hostname)
        hostname, err := os.Hostname()
        if err != nil {
                log.Fatalf("could not detect hostname: %v", err)
        }
        return &pb.HelloReply{Message: "Hello " + in.Name, Hostname: hostname, NodeName: nodeName, PodName: podName}, nil
}

func main() {
        lis, err := net.Listen("tcp", host+":"+port)
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
        }
        s := grpc.NewServer()
        pb.RegisterGreeterServer(s, &server{})
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
		os.Exit(2)
        }
}
