package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type ErrorStruct struct {
	isError     bool
	ErrorString string
}

func (calendar *CalendarRequest) validate() (err ErrorStruct) {

	if calendar.Name == nil {
		err.isError = true
		err.ErrorString = "Name cannot be empty"
		return err
	}

	if calendar.Active == nil {
		err.isError = true
		err.ErrorString = "Active field cannot be empty"
		return err
	}

	if calendar.Color == nil {
		err.isError = true
		err.ErrorString = "Color field cannot be empty"
		return err
	}

	if calendar.Overlap == nil {
		err.isError = true
		err.ErrorString = "Overlap field cannot be empty"
		return err
	}

	if calendar.Attributes == nil {
		err.isError = true
		err.ErrorString = "Attributes field cannot be empty"
		return err
	}

	err.isError = false
	err.ErrorString = ""
	return err

}

func CreateCalendar(w http.ResponseWriter, r *http.Request) {

	// dbConn := ConnectDb()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var calendar CalendarRequest
	fmt.Println("calnedar is ", calendar)
	err := json.Unmarshal(reqBody, &calendar)
	fmt.Println("Error is ", err)
	fmt.Println("Calendar is ", (calendar))
	fmt.Println("And is ", calendar.Name)
	rand.Seed(time.Now().UTC().UnixNano())
	validationErr := calendar.validate()
	if validationErr.isError == true {
		fmt.Println(validationErr.ErrorString)
		w.WriteHeader(400)
		return
	}
	// id := strconv.Itoa(randomInt(10000, 99999))
	// name := calendar.Name
	// active := calendar.Active
	// color := calendar.Color
	// overlap := (calendar.Overlap)
	// attributes := calendar.Attributes
	// location := "Dehradun"

	// stmtIns, err := dbConn.Prepare("INSERT INTO calendar(id, name, active, color, overlap, attributes, location) VALUES (?, ?, ?, ?, ?, ?, ?) ")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer stmtIns.Close()

	// _, err = stmtIns.Exec(id, name, active, color, overlap, attributes, location)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// w.WriteHeader(201)
	// json.NewEncoder(w).Encode(CalendarRequest(calendar))
	fmt.Println("Data is ", calendar)
}
