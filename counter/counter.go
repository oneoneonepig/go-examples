package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
	KeyName string `json:"key"`
	Value int64 `json:"value"`
}

func Homepage(w http.ResponseWriter, r *http.Request){
	title := "Homepage for counting"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, title)
}

func OptionForCORS(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.WriteHeader(http.StatusOK)
}

// Get a single key's value
func GetKey(w http.ResponseWriter, r *http.Request){
	pathParams := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	keyName := pathParams["key"]
	result := rdb.Get(keyName)
	// Bad
	if result.Err() != nil {
		response := Response{
			Message: "key retrieve failed",
			KeyName: keyName,
			Value:   0,
		}
		json.NewEncoder(w).Encode(response)
	// Good
	} else {
		value, _ := strconv.ParseInt(result.Val(), 10, 64)
		response := Response{
			Message: "key retrieved",
			KeyName: keyName,
			Value:   value,
		}
		json.NewEncoder(w).Encode(response)
	}
}

// Set a single key's value
func SetKey(w http.ResponseWriter, r *http.Request){
	pathParams := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	keyName := pathParams["key"]
	value := pathParams["value"]
	result := rdb.Set(keyName, value, 0)
	// Bad
	if result.Err() != nil {
		response := Response{
			Message: result.Err().Error(),
		}
		json.NewEncoder(w).Encode(response)
	// Good
	} else {
		value, _ := strconv.ParseInt(value, 10, 64)
		response := Response{
			Message: "key set",
			KeyName: keyName,
			Value: value,
		}
		json.NewEncoder(w).Encode(response)
	}

}

// Delete a single key
func DeleteKey(w http.ResponseWriter, r *http.Request){
	pathParams := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	keyName := pathParams["key"]
	rdb.Del(keyName)
	response := Response{
		Message: "key deleted",
		KeyName: keyName,
		Value: 0,
	}
	json.NewEncoder(w).Encode(response)
}

// Get all keys' names
func GetAllKeys(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var cursor uint64
	var n int
	var allKeys []string
	for {
		var keys []string
		var err error
		keys, cursor, err = rdb.Scan(cursor, "*", 10).Result()
		if err != nil {
			panic(err)
		}
		n += len(keys)
		allKeys = append(allKeys, keys...)
		if cursor == 0 {
			break
		}
	}
	response := Response{
		Message: "",
		KeyName: "",
		Value:   0,
	}
	for _, k := range allKeys {
		v, _ := rdb.Get(k).Result()
		response.Message = response.Message + k + " = " + v + "<br>"
	}

	json.NewEncoder(w).Encode(response)
}

// Delete all keys
func DeleteAllKeys(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	rdb.FlushAll()
	response := Response{
		Message: "db flushed",
		KeyName: "",
		Value:   0,
	}
	json.NewEncoder(w).Encode(response)
}
