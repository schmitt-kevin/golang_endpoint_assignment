package main

import (
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"net/http"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	// "fmt"

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

// GetPersonEndpoint get single person
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	params := mux.Vars(req)
	for _, item := range customers {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Customer{})
}

// GetPeopleEndpoint get all people
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(customers)
}

// DeletePersonEndpoint delete single person
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	params := mux.Vars(req)
	for i, item := range customers {
		if item.ID == params["id"] {
			customers = append(customers[:i], customers[i+1])
			break
		}
	}
	json.NewEncoder(w).Encode(customers)
}

// CreatePersonEndpoint creates a single person
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	var customer Customer
	lastIdx := len(customers) - 1
	lastCustomer := customers[lastIdx]
	_ = json.NewDecoder(req.Body).Decode(&customer)
	lastID, err := strconv.Atoi(lastCustomer.ID)
	if err != nil {
		log.Fatalf("Cannot convert str to int: %s", err.Error())
	}
	customer.ID = strconv.Itoa(lastID + 1)
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer)
}

// EditPersonEndpoint edit an already existing person
func EditPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	params := mux.Vars(req)
	var customer Customer
	for i, item := range customers {
		if item.ID == params["id"] {
			customer.ID = params["id"]
			_ = json.NewDecoder(req.Body).Decode(&customer)
			customers[i] = customer
			break
		}
	}
	json.NewEncoder(w).Encode(customer)
}

// DownloadAddressBook export a csv file to project folder
func DownloadAddressBook(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	rows := GetRowsFromCustomersJson()
	f, err := os.Create("ExportAddressBook.csv")
	if err != nil {
		log.Fatalf("Cannot create a downloadable address book: %s", err.Error())
	}

	defer func() {
		e := f.Close()
		if e != nil {
		log.Fatalf("Cannot close created file: %s", err.Error())
		}
	}()

	d := csv.NewWriter(f)
	err = d.WriteAll(rows)
}

//GetRowsFromCustomersJson reads the slice of current customers and placing into a [][]string
func GetRowsFromCustomersJson() [][]string {
	var rows [][]string
	for i, item := range customers {
		var row []string
		if i == 0 {
			row = append(row, "First Name")
			row = append(row, "Last Name")
			row = append(row, "Phone")
			row = append(row, "Email")
			rows = append(rows, row)
		}
		row = []string{}
		row = append(row, item.FirstName)
		row = append(row, item.LastName)
		row = append(row, item.Phone)
		row = append(row, item.Email)
		rows = append(rows, row)		
	}

	return rows
}

// ImportAddressBook import values from csv file in project folder
func ImportAddressBook(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	rows := readCSV("ImportAddressBook.csv")
	var customer Customer
	for _, row := range rows {
		if strings.ToLower(row[0]) != "first name"{
			customer.FirstName = row[0]
			customer.LastName = row[1]
			customer.Phone = row[2]
			customer.Email = row[3]
			customer.ID = strconv.Itoa(len(customers) + 1)
			customers = append(customers, customer)
			customer = Customer{}
		}
	}
	json.NewEncoder(w).Encode(customers)
}

// readCSV reads the file needing to be imported
func readCSV(name string) [][]string {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot open csv file:%s", err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','

	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalf("Cannot read csv file:%s", err.Error())
	}

	return rows
}

func main() {
	router := mux.NewRouter()

	customers = append(customers, Customer{ID: "1", FirstName: "Kevin", LastName: "Schmitt", Phone: "817-703-9740", Email: "schmittkw@yahoo.com"})
	customers = append(customers, Customer{ID: "2", FirstName: "Wendy", LastName: "Quest", Phone: "817-303-2195", Email: "wendy_quest@yahoo.com"})

	router.HandleFunc("/customer", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/customer/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/customer", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/customer/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/customer/{id}", EditPersonEndpoint).Methods("PUT")
	router.HandleFunc("/download", DownloadAddressBook).Methods("GET")
	router.HandleFunc("/import", ImportAddressBook).Methods("GET")

	http.ListenAndServe(":1337", router)

}