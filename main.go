package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Lucky!")
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Starting the server on port 10000...")
	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal(err)
	}
}
