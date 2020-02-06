package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEventRange(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	enddate, _ := r.URL.Query()["enddate"]
	startdate, _ := r.URL.Query()["startdate"]
	calendarId := params["calendarId"]
	var event []EventRequest
	dbConn := ConnectDb()
	var err error
	if startdate != nil && enddate != nil {
		err = dbConn.Select(&event, "SELECT * FROM events where calendarid=? AND start_dt >= ? AND end_dt <= ?", calendarId, startdate[0], enddate[0])
	} else if startdate != nil {
		err = dbConn.Select(&event, "SELECT * FROM events where calendarid=? AND start_dt >= ?", calendarId, startdate[0])
	} else if enddate != nil {
		err = dbConn.Select(&event, "SELECT * FROM events where calendarid=? AND end_dt <= ?", calendarId, enddate[0])
	} else {
		err = dbConn.Select(&event, "SELECT * FROM events where calendarid=? ", calendarId)
	}
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(403)
		return
	}
	json.NewEncoder(w).Encode([]EventRequest(event))
}
