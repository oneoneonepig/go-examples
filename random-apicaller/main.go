package main

import (
	"log"
	flag "github.com/spf13/pflag"
	"math/rand"
	"net/http"
	"time"
)

var (
	url  = flag.StringP("url", "u", "http://localhost", "url endpoint")
	prob = flag.IntP("probability", "p", 100, "one in how much that the api would be called in a second")
)

func main() {
	flag.Parse()
	client := http.Client{
		Timeout: time.Duration(1 * time.Second),
	}
	for {
		if rand.Int()%*prob == 0 {
			log.Println("hooray! calling endpoint: " + *url)
			_, err := client.Get(*url)
			if err != nil {
				log.Println("sadly, cannot reach endpoint: " + *url)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
