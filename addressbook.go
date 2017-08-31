//Address Book
//by Kirby Flake

package main

import (
	"encoding/json"
	"encoding/csv"
	"fmt"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

//data: First Name, Last Name, Email Address, and Phone Number
//list
//add
//remove
//update

type Record struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

var addressRecord []Record

//Get All Records from addressbook.csv
func importRecords() {
	file, err := os.Open("addressbook.csv")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("New addressbook")
		}
		//return
	}
	defer file.Close()
	if err == nil {
		reader := csv.NewReader(file)
		record, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error", err)
		}

		for value := range record { // for i:=0; i<len(records)
			addressRecord = append(addressRecord, Record{
				ID:        record[value][0],
				FirstName: record[value][1],
				LastName:  record[value][2],
				Email:     record[value][3],
				Phone:     record[value][4],
			})
		}
	}
}

//Writes All Records from addressbook.csv
func exportRecords() {
	file, err := os.Create("addressbook.csv")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for value := range addressRecord {
		record := []string{
			addressRecord[value].ID,
			addressRecord[value].FirstName,
			addressRecord[value].LastName,
			addressRecord[value].Email,
			addressRecord[value].Phone,
		}
		err := writer.Write(record)

		if err != nil {
			fmt.Println("Error", err)
		}
	}
}

//Returns a JSON of all records in the addressbook
func listRecords(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(addressRecord)
}

//Returns a JSON of record of the ID requested
func listRecord(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range addressRecord {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Record{})
}

//Returns a JSON of record of the ID requested
func addRecord(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var record Record
	_ = json.NewDecoder(req.Body).Decode(&record)
	record.ID = params["id"]
	addressRecord = append(addressRecord, record)
	json.NewEncoder(w).Encode(addressRecord)
	exportRecords()
}

//Deletes record with ID requested
func deleteRecord(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range addressRecord {
		if item.ID == params["id"] {
			addressRecord = append(addressRecord[:index], addressRecord[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(addressRecord)
	exportRecords()
}

//Modifies record with ID requested
func modifyRecord(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var record Record
	for index, item := range addressRecord {
		if item.ID == params["id"] {
			_ = json.NewDecoder(req.Body).Decode(&record)
			addressRecord[index].FirstName = record.FirstName
			addressRecord[index].LastName = record.LastName
			addressRecord[index].Email = record.Email
			addressRecord[index].Phone = record.Phone
			break
		}
	}
	json.NewEncoder(w).Encode(addressRecord)
	exportRecords()
}

func main() {
	importRecords()
	//addressRecord = append(addressRecord, Record{ID: "1", FirstName: "Joseph", LastName: "Smith", Email: "test@test.biz", Phone: "555-555-5555" })
	//addressRecord = append(addressRecord, Record{ID: "2", FirstName: "Josepha", LastName: "Smith", Email: "test2@test.biz", Phone: "554-555-5555"})
	//exportRecords()
	router := mux.NewRouter()
	router.HandleFunc("/record", listRecords).Methods("GET") //Lists all records
	router.HandleFunc("/record/{id}", listRecord).Methods("GET") //List a specific record
	router.HandleFunc("/record/{id}", addRecord).Methods("POST") 	//Adds new record
	router.HandleFunc("/record/{id}", deleteRecord).Methods("DELETE") //Deletes the record
	router.HandleFunc("/record/{id}", modifyRecord).Methods("PATCH") //Modifies the record
	//PUT	Not used
	log.Fatal(http.ListenAndServe(":8081", router))

}
