package main

import (
	"github.com/gorilla/mux"
	"encoding/json"
	//"fmt"
	"log"
	"net/http"

)

//Customer type
type Customer struct {
	ID string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

var customers []Customer

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range customers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Customer{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(customers)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, item := range customers {
		if item.ID == params["id"] {
			customers = append(customers[:i], customers[i+1])
			break
		}
	}
	json.NewEncoder(w).Encode(customers)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
params := mux.Vars(req)
	var customer Customer
	_ = json.NewDecoder(req.Body).Decode(&customer)
	customer.ID = params["id"]
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customers)
}

func EditPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for i, item := range customers {
		if item.ID == params["id"] {
			var customer Customer
			customer.ID = params["id"]
			_ = json.NewDecoder(req.Body).Decode(&customer)
			customers[i] = customer
			break
		}
	}
	json.NewEncoder(w).Encode(customers)
}

func main() {
	router := mux.NewRouter()
	customers = append(customers, Customer{ID: "1", FirstName: "Kevin", LastName: "Schmitt", Phone: "817-703-9740", Email: "schmittkw@yahoo.com"})
	customers = append(customers, Customer{ID: "2", FirstName: "Wendy", LastName: "Quest", Phone: "817-303-2195", Email: "wendy_quest@yahoo.com"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/people/{id}", EditPersonEndpoint).Methods("PUT")
	log.Fatal(http.ListenAndServe(":1337", router))

}