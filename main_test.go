package main

import (
	"testing"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
)

func TestCustomerAPI(t *testing.T) {
	customers = append(customers, Customer{ID: "1", FirstName: "Kevin", LastName: "Schmitt", Phone: "817-703-9740", Email: "schmittkw@yahoo.com"})
	customers = append(customers, Customer{ID: "2", FirstName: "Wendy", LastName: "Quest", Phone: "817-303-2195", Email: "wendy_quest@yahoo.com"})
	testGetAll(t)
	testGetPerson(t)
	testDeletePerson(t)
	testCreatePerson(t)
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/customer", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/customer/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/customer", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/customer/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/customer/{id}", EditPersonEndpoint).Methods("PUT")
	router.HandleFunc("/download", DownloadAddressBook).Methods("GET")
	router.HandleFunc("/import", ImportAddressBook).Methods("GET")
	return router
}

func testGetAll(t *testing.T) {
	request, _ := http.NewRequest(/*method*/ "GET", /*url*/ "/customer", /*body*/ nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, /*expected*/ 200, /*received*/ response.Code, /*msgAndArgs*/ "OK response is expected")
	assert.Equal(t, /*expected*/ `[{"id":"1","first_name":"Kevin","last_name":"Schmitt","email":"schmittkw@yahoo.com","phone":"817-703-9740"},{"id":"2","first_name":"Wendy","last_name":"Quest","email":"wendy_quest@yahoo.com","phone":"817-303-2195"}]`+"\n", /*received*/ response.Body.String(), /*msgAndArgs*/ "OK response is expected")
}

func testGetPerson(t *testing.T) {
	request, _ := http.NewRequest(/*method*/ "GET", /*url*/ "/customer/1", /*body*/ nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, /*expected*/ 200, /*received*/ response.Code, /*msgAndArgs*/ "OK response is expected")
	responseStr := response.Body.String()
	if !strings.Contains(responseStr, "Kevin") {
		t.Fatalf("Did not receive the correct response from testGetPerson")
	}
}

func testDeletePerson(t *testing.T) {
	request, _ := http.NewRequest(/*method*/ "DELETE", /*url*/ "/customer/1", /*body*/ nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, /*expected*/ 200, /*received*/ response.Code, /*msgAndArgs*/ "OK response is expected")
	responseStr := response.Body.String()
	if strings.Contains(responseStr, "Kevin") {
		t.Fatalf("Did not receive the correct response from testDeletePerson")
	}
}

func testCreatePerson(t *testing.T) {
	var customer Customer
	customer.FirstName = "Daniel"
	customer.LastName = "George"
	customer.Phone = "932-943-2758"
	customer.Email = "dannyg@email.com"
	lastIdx := len(customers) - 1
	lastCustomer := customers[lastIdx]
	lastID, err := strconv.Atoi(lastCustomer.ID)
	if err != nil {
		t.Fatalf("Cannot convert str to int: %s", err.Error())
	}
	customer.ID = strconv.Itoa(lastID + 1)
	customers = append(customers, customer)
	if customer.ID != "3" {
		t.Fatalf("Had trouble adding new customer into customers slice")
	}	
}

