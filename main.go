package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Book_Name   string
	Publication string
	Chapter     int
	Price       float64
}

var Product []Book

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpont hit:homepage")
	fmt.Fprintf(w, "Hey there lucky")
}

func returnAllProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProduct")
	json.NewEncoder(w).Encode(Product)
}

func handleRequests() {
	http.HandleFunc("/Product", returnAllProduct)
	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:10000", nil)

}

func main() {
	Product = []Book{
		{Book_Name: "Science", Publication: "Mittal", Chapter: 15, Price: 390.00},
		{Book_Name: "Mathematics", Publication: "Oxford", Chapter: 12, Price: 450.00},
		{Book_Name: "History", Publication: "Penguin", Chapter: 20, Price: 550.00},
		{Book_Name: "Hindi", Publication: "Mittal", Chapter: 10, Price: 350.00},
		{Book_Name: "Economy", Publication: "Oxford", Chapter: 17, Price: 300.00},
	}

	handleRequests()

}
