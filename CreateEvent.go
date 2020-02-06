package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (event *EventRequest) validateEvent() (err ErrorStruct) {

	if event.Start_dt == nil {
		err.isError = true
		err.ErrorString = "Start date cannot be empty"
		return err
	}

	if event.End_dt == nil {
		err.isError = true
		err.ErrorString = "End date cannot be empty"
		return err
	}

	if event.All_day == nil {
		err.isError = true
		err.ErrorString = "All day field cannot be empty"
		return err
	}

	if event.Title == nil {
		err.isError = true
		err.ErrorString = "Title cannot be empty"
		return err
	}

	if event.Who == nil {
		err.isError = true
		err.ErrorString = "Who field cannot be empty"
		return err
	}

	if event.Location == nil {
		err.isError = true
		err.ErrorString = "Location field cannot be empty"
		return err
	}

	if event.Notes == nil {
		err.isError = true
		err.ErrorString = "Notes field cannot be empty"
		return err
	}

	err.isError = false
	err.ErrorString = ""
	return err
}

//create table events (id varchar(20), subCalendarId varchar(20), calendarId varchar(20), version varchar(20), creation_dt datetime, update_dt datetime, start_dt datetime, end_dt datetime, allday bool, title varchar(100), who varchar(200), location varchar(200), notes varchar(400));

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	calendarId := params["calendarId"]

	dbConn := ConnectDb()
	reqBody, _ := ioutil.ReadAll(r.Body)
	var event EventRequest

	_ = json.Unmarshal(reqBody, &event)
	validationErr := event.validateEvent()
	stmtIns, err := dbConn.Prepare("INSERT INTO events(id, subcalendarid, calendarid, version, creation_dt, update_dt, start_dt, end_dt, all_day, title, who, location, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	if validationErr.isError == true {
		fmt.Println(validationErr.ErrorString)
		w.WriteHeader(400)
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	id := strconv.Itoa(randomInt(10000, 99999))
	subCalendarId := strconv.Itoa(randomInt(10000, 99999))
	version := strconv.Itoa(randomInt(10000, 99999))

	currentDate := time.Now()
	creation_dt := currentDate.Format("2006-01-02 15:04:05")
	update_dt := currentDate.Format("2006-01-02 15:04:05")

	_, err = stmtIns.Exec(id, subCalendarId, calendarId, version, creation_dt, update_dt, event.Start_dt, event.End_dt, event.All_day, event.Title, event.Who, event.Location, event.Notes)

	if err != nil {
		panic(err.Error())
	}
	w.WriteHeader(201)
	var eventResponse EventRequest

	row := dbConn.QueryRowx("SELECT * FROM events where id =?", id)
	err = row.StructScan(&eventResponse)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(EventRequest(eventResponse))
}
