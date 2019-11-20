package main

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Counter int64 `json:"counter"`
}

/*
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

}

 */

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

func Retrieve(w http.ResponseWriter, r *http.Request){
	// enableCors(&w)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result, err := rdb.Get("counter").Int64()
	if err != nil {
		response := Response{
			Message: "key not found",
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response := Response{
			Counter: result,
			Message: "value retrieved",
		}
		json.NewEncoder(w).Encode(response)
	}
}

func Increase(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	result, err := rdb.Incr("counter").Result()
	if err != nil {
		panic(err)
	}
	response := Response{
		Counter: result,
		Message: "value increased",
	}
	json.NewEncoder(w).Encode(response)
}

func Delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	rdb.Del("counter")
	response := Response{
		Message: "key deleted",
	}
	json.NewEncoder(w).Encode(response)
}

func notFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	//w.Write([]byte(`method not implemented`))
	w.Write([]byte(r.Method))
}
