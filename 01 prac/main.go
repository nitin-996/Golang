package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)




type Product struct {

	Id int
	Name string
	Quantity int
	Price float64
}


var Products []Product

func homepage(w http.ResponseWriter , r *http.Request){
	fmt.Fprint(w, "welcome to the home page")
	fmt.Println("Endpoint hit: HomePage")
}


func returnAllProducts(w http.ResponseWriter, r *http.Request){
log.Println("Endpoint Hit: returnAllProducts")
json.NewEncoder(w).Encode(Products)
}
func handlerRequest(){

	http.HandleFunc("/products", returnAllProducts )
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":3000", nil)
}


func main(){
	
	Products = []Product{
		{Id:1, Name:"nemo",Quantity: 10,Price: 100},
		{Id:2, Name:"shark",Quantity: 10,Price: 200},

	}
	
}