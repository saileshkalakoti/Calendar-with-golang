package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func PatchCalendar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	dbConn := ConnectDb()

	reqBody, _ := ioutil.ReadAll(r.Body)

	var calendar CalendarRequest

	err := json.Unmarshal(reqBody, &calendar)

	name := calendar.Name
	active := calendar.Active
	color := calendar.Color
	overlap := (calendar.Overlap)
	attributes := calendar.Attributes

	currentDate := time.Now()
	update_dt := currentDate.Format("2006-01-02 15:04:05")

	res := dbConn.MustExec("UPDATE calendar SET name=?, active=?, color=?, overlap=?, attributes=?, update_dt=? WHERE id=?", name, active, color, overlap, attributes, update_dt, id)
	count, err := res.RowsAffected()
	if err != nil {
		w.WriteHeader(500)
	}
	if count == 1 {
		row := dbConn.QueryRowx("SELECT * FROM calendar where id =?", id)
		_ = row.StructScan(&calendar)

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(CalendarRequest(calendar))
	}
}
