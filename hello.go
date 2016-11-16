package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/h", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})
	mux.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "byebye")
	})
	mux.HandleFunc("/hello", sayhello)
	http.ListenAndServe(":8080", mux)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}
