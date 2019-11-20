package main

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	rdb           *redis.Client
	redisAddr     = "10.20.131.50"
	redisPort     = "6379"
	redisPassword = ""
)

func main() {
	// Create router
	r := mux.NewRouter()
	
	r.HandleFunc("/", Homepage).Methods(http.MethodGet)

	rc := r.PathPrefix("/counter").Subrouter()
	rc.HandleFunc("", Retrieve).Methods(http.MethodGet)
	rc.HandleFunc("", Increase).Methods(http.MethodPost)
	rc.HandleFunc("", Delete).Methods(http.MethodDelete)
	rc.HandleFunc("", OptionForCORS).Methods(http.MethodOptions)

	/*
	r.HandleFunc("/counter", Retrieve).Methods(http.MethodGet)
	r.HandleFunc("/counter", Increase).Methods(http.MethodPost)
	r.HandleFunc("/counter", Delete).Methods(http.MethodDelete)
	r.HandleFunc("/counter", OptionForCORS).Methods(http.MethodOptions)
	*/

	// Create redis connection
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr + ":" + redisPort,
		Password: redisPassword,
		DB:       0,
	})

	// Test redis connection
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}

	// Start serving HTTP
	log.Fatal(http.ListenAndServe(":8080", r))
}
