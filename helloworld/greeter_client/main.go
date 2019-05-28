package main

import (
        "context"
        "flag"
        "log"
        "os"
        "time"

	"encoding/json"
	"net/http"

        "google.golang.org/grpc"
        pb "github.com/oneoneonepig/go-examples/helloworld/helloworld"
)

type FrontendResponse struct {
	NodeName	string `json:"nodeName"`
	PodName		string `json:"podName"`
}

var (
        host		string
        port		string
	listenPort	string
        message		string
)

func init() {
        flag.StringVar(&host, "host", "0.0.0.0", "Target address")
        flag.StringVar(&port, "port", "3000", "Target port")
        flag.StringVar(&listenPort, "listen", "8080", "Listening port")
        flag.StringVar(&message, "message", "world", "Sending message")
        flag.Parse()
}

func main() {

	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	        c := pb.NewGreeterClient(conn)

	        name := message
	        hostname, err := os.Hostname()
	        if err != nil {
	                log.Fatalf("could not detect hostname: %v", err)
	        }

	        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	        defer cancel()
	        rr, err := c.SayHello(ctx, &pb.HelloRequest{Name: name, Hostname: hostname})
	        if err != nil {
	                log.Fatalf("could not greet: %v", err)
	        }
		log.Printf("Reply: %s, from %s, node: %s, pod: %s", rr.Message, rr.Hostname, rr.NodeName, rr.PodName)

		var fr FrontendResponse
		fr.NodeName = rr.NodeName
		fr.PodName = rr.PodName
		json.NewEncoder(w).Encode(fr)
	})

	log.Fatal(http.ListenAndServe(":"+listenPort, nil))
}
