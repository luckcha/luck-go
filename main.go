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

type Dealer struct {
	ID       int
	Name     string
	Location string
	sale     float64
}

var Product []Book
var Production []Dealer

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpont hit:homepage")
	fmt.Fprintf(w, "Hey there lucky")
}

func returnAllProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnAllProduct")
	json.NewEncoder(w).Encode(Product)
}
func returnAllProduction(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit:returnAllProduction")
	json.NewEncoder(w).Encode(Production)
}

func handleRequests() {
	http.HandleFunc("/Product", returnAllProduct)
	http.HandleFunc("/Production", returnAllProduction)
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
	Production = []Dealer{
		{ID: 101, Name: "Jack", Location: "Minnisota", sale: 455658.5},
		{ID: 101, Name: "Luther", Location: "New York", sale: 358858.8},
		{ID: 101, Name: "Garry", Location: "London", sale: 625895.8},
		{ID: 101, Name: "Adrian", Location: "Paris", sale: 9427158.5},
		{ID: 101, Name: "Mickel", Location: "Washington", sale: 575754.4},
	}

	handleRequests()

}
