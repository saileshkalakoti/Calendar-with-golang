package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Calendar struct {
	Active      bool      `json:"active"`
	Color       int       `json:"color"`
	Overlap     bool      `json:"overlap"`
	Attributes  []string  `json:"attributes"`
	Creation_dt time.Time `json:"creation_dt"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	Update_dt   time.Time `json:"update_dt"`
	Id          string    `json:"id"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("THIS IS GOOD")
	params := mux.Vars(r)
	fmt.Println("Params are ", params)
}

func createCalendar(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var calendar Calendar
	err := json.Unmarshal(reqBody, &calendar)
	fmt.Println("Error is ", err)
	fmt.Println("Calendar is ", calendar)

	json.NewEncoder(w).Encode(Calendar(calendar))
	fmt.Println("Data is ", calendar)
}

func main() {
	currentTime := time.Now()
	fmt.Println("Current date is ", currentTime)
	r := mux.NewRouter()
	r.HandleFunc("/articles/{id}", handler).Methods("GET")
	r.HandleFunc("/api/calendar", createCalendar).Methods("POST")
	http.ListenAndServe(":8080", r)
}
