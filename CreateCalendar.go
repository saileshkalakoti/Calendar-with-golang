package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)


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

func getImageUrl() string {
	resp, _ := http.Get("https://source.unsplash.com/random")
	return resp.Request.URL.String()
}

func CreateCalendar(w http.ResponseWriter, r *http.Request) {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan string)
	go func() { ch <- getImageUrl() }()
	dbConn := ConnectDb()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var calendar CalendarRequest

	err := json.Unmarshal(reqBody, &calendar)

	rand.Seed(time.Now().UTC().UnixNano())
	validationErr := calendar.validate()
	if validationErr.isError == true {
		fmt.Println(validationErr.ErrorString)
		w.WriteHeader(400)
		return
	}
	id := strconv.Itoa(randomInt(10000, 99999))
	name := calendar.Name
	active := calendar.Active
	color := calendar.Color
	overlap := (calendar.Overlap)
	attributes := calendar.Attributes
	location := "Dehradun"

	currentDate := time.Now()
	creation_dt := currentDate.Format("2006-01-02 15:04:05")
	update_dt := currentDate.Format("2006-01-02 15:04:05")
	stmtIns, err := dbConn.Prepare("INSERT INTO calendar(id, name, active, color, overlap, attributes, location, creation_dt, update_dt, image_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(id, name, active, color, overlap, attributes, location, creation_dt, update_dt, <-ch)
	wg.Done()
	if err != nil {
		panic(err.Error())
	}
	w.WriteHeader(201)
	var createdCalendar CalendarRequest
	row := dbConn.QueryRowx("SELECT * FROM calendar where id =?", id)
	err = row.StructScan(&createdCalendar)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(CalendarRequest(createdCalendar))
}
