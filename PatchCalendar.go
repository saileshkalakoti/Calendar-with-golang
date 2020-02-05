package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func PatchCalendar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	dbConn := ConnectDb()

	reqBody, _ := ioutil.ReadAll(r.Body)
	var calendar CalendarRequest
	err := json.Unmarshal(reqBody, &calendar)
	fmt.Println("Error is ", err)
	fmt.Println("Calendar is ", calendar)
	name := calendar.Name
	active := calendar.Active
	color := calendar.Color
	overlap := (calendar.Overlap)
	attributes := calendar.Attributes
	// location := "Dehradun"
	res := dbConn.MustExec("UPDATE calendar SET name=?, active=?, color=?, overlap=?, attributes=? WHERE id=?", name, active, color, overlap, attributes, id)
	count, err := res.RowsAffected()
	if err != nil {
		panic(err.Error)
	}
	if count == 1 {
		row := dbConn.QueryRowx("SELECT name, id, active, location, color FROM calendar where id =?", id)
		// fmt.Println(row)
		_ = row.StructScan(&calendar)

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(CalendarRequest(calendar))
	}
}
