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

// Package main implements a client for Greeter service.
package main

import (
        "context"
        "flag"
        "log"
        "os"
        "time"

        "google.golang.org/grpc"
        pb "github.com/oneoneonepig/go-examples/helloworld/helloworld"
)

/*
const (
         address     = "localhost:32123"
        defaultName = "world"
)
*/

var (
        host    string
        port    string
        message string
)

// Initialize listen port and address
func init() {
        flag.StringVar(&host, "host", "0.0.0.0", "help message for host")
        flag.StringVar(&port, "port", "3000", "help message for port")
        flag.StringVar(&message, "message", "world", "help message for message")
}

func main() {
        flag.Parse()

        // Set up a connection to the server.
        conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        c := pb.NewGreeterClient(conn)

        name := message
        hostname, err := os.Hostname()
        if err != nil {
                log.Fatalf("could not detect hostname: %v", err)
        }

        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name, Hostname: hostname})
        if err != nil {
                log.Fatalf("could not greet: %v", err)
        }
        log.Printf("Reply: %s from %s, node: %s, pod: %s", r.Message, r.Hostname, r.NodeName, r.PodName)
}
