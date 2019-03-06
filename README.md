# golang_endpoint_assignment

To save some time I did not include a DB

## Needed before starting server
### -router
### go get github.com/gorilla/mux 
### -testing
### go get github.com/stretchr/testify/assert

## Start server
### ./golang_endpoint_assignment OR go run main.go

## Using API
### open postman

## Get all
### -GET
### localhost:1337/customer

## Get single
### -GET
### localhost:1337/customer/{id}

## Edit
### -PUT
### localhost:1337/customer/{id}

## Delete
### -DELETE
### localhost:1337/customer/{id}

## Create
### -POST
### localhost:1337/customer
### -pass a json object matching the customer struct
{
    "first_name": "",
    "last_name": "",
    "phone": "",
    "email": ""
}

## Export CSV
### localhost:1337/download

### -will download to same folder you run main.go from

## Import
### -will grab a csv file by the name "ImportAddressBook.csv"
### -this is located in the same folder as main.go

## Testing
### go test
### -will run through tests in main_test.go
