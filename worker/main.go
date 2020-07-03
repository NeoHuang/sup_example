package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello this is worker\n")
}

func main() {
	http.HandleFunc("/", hello)
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT is not set")
	}

	log.Printf("Starting server at port:%s", port)
	http.ListenAndServe(":"+port, nil)
}
