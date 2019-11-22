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
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.HandleFunc("/keys", GetAllKeys).Methods(http.MethodGet)
	r.HandleFunc("/keys", DeleteAllKeys).Methods(http.MethodDelete)
	r.HandleFunc("/key/{key}", GetKey).Methods(http.MethodGet)
	r.HandleFunc("/key/{key}/{value:[0-9]+}", SetKey).Methods(http.MethodPost)
	r.HandleFunc("/key/{key}", DeleteKey).Methods(http.MethodDelete)
	r.HandleFunc("/", OptionForCORS).Methods(http.MethodOptions)

	/*
	r.HandleFunc("/counter", Retrieve).Methods(http.MethodGet)
	r.HandleFunc("/counter", Increase).Methods(http.MethodPost)
	r.HandleFunc("/counter", Delete).Methods(http.MethodDelete)

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
