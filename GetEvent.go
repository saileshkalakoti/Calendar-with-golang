package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	calendarId := params["calendarId"]
	eventId := params["eventId"]

	dbConn := ConnectDb()
	var event EventRequest
	row := dbConn.QueryRowx("SELECT * FROM events where id =? AND calendarid=? ", eventId, calendarId)
	err := row.StructScan(&event)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(403)
		return
	}
	json.NewEncoder(w).Encode(EventRequest(event))

}
