package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCalendar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	dbConn := ConnectDb()
	var calendar CalendarRequest
	row := dbConn.QueryRowx("SELECT * FROM calendar where id =?", id)
	err := row.StructScan(&calendar)
	if err != nil {
		w.WriteHeader(403)
		return
	}
	json.NewEncoder(w).Encode(CalendarRequest(calendar))
}
