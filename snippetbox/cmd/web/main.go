package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view",snippetView)
    mux.HandleFunc("/snippet/create",snippetCreate)

    log.Print("Starting a Server on Port 4000")
    err := http.ListenAndServe(":8000", mux)
    log.Fatal(err)
}
