package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view",snippetView)
    mux.HandleFunc("/snippet/create",snippetCreate)

    log.Print("Starting a Server on Port 8000")
    err := http.ListenAndServe(":8000", mux)
    log.Fatal(err)
}
