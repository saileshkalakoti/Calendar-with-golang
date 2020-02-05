package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteCalendar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	dbConn := ConnectDb()
	res := dbConn.MustExec("DELETE FROM calendar WHERE id = ?", id)
	count, err := res.RowsAffected()

	if err != nil {
		w.WriteHeader(500)
		panic(err.Error)
	}
	if count == 1 {
		w.WriteHeader(204)
	}
}
