package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func CreateCalendar(w http.ResponseWriter, r *http.Request) {

	dbConn := ConnectDb()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var calendar CalendarRequest
	err := json.Unmarshal(reqBody, &calendar)
	fmt.Println("Error is ", err)
	fmt.Println("Calendar is ", calendar)
	rand.Seed(time.Now().UTC().UnixNano())
	id := strconv.Itoa(randomInt(10000, 99999))
	name := calendar.Name
	active := calendar.Active
	color := calendar.Color
	overlap := (calendar.Overlap)
	attributes := calendar.Attributes
	location := "Dehradun"

	stmtIns, err := dbConn.Prepare("INSERT INTO calendar(id, name, active, color, overlap, attributes, location) VALUES (?, ?, ?, ?, ?, ?, ?) ")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(id, name, active, color, overlap, attributes, location)
	if err != nil {
		panic(err.Error())
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(CalendarRequest(calendar))
	fmt.Println("Data is ", calendar)
}
