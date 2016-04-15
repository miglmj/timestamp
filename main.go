package main

import (
	"net/http"
)

type timestamp struct {
	Human string
	Unix  string
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}
