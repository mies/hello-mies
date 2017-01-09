package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello mies!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloHandler)
	http.Handle("/", r)
}
