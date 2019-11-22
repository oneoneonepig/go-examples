# Counter

Operate Redis key/value on web, written by Go

## API

```
r.HandleFunc("/", Homepage).Methods(http.MethodGet)
r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
r.HandleFunc("/keys", GetAllKeys).Methods(http.MethodGet)
r.HandleFunc("/keys", DeleteAllKeys).Methods(http.MethodDelete)
r.HandleFunc("/key/{key}", GetKey).Methods(http.MethodGet)
r.HandleFunc("/key/{key}/{value}", SetKey).Methods(http.MethodPost)
r.HandleFunc("/key/{key}", DeleteKey).Methods(http.MethodDelete)
r.HandleFunc("/", OptionForCORS).Methods(http.MethodOptions)
```
