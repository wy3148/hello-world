package main

import (
	"log"
	"net/http"
	"os"
)

func defaultView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3333"
	}

	http.HandleFunc("/", defaultView)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
