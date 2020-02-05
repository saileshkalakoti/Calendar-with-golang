package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Calendar struct {
	Creation_dt time.Time `json:"creation_dt"`
	Location    string    `json:"location"`
	Update_dt   time.Time `json:"update_dt"`
	Id          string    `json:"id"`
	CalendarRequest
}

type CalendarRequest struct {
	Name       *string `json:"name"`
	Active     *bool   `json:"active"`
	Color      *int    `json:"color"`
	Overlap    *bool   `json:"overlap"`
	Attributes *string `json:"attributes"`
}

// insert into calendar(id, name, active, color, overlap, attributes, location) values ('123', 'sailesh', 1, 4, 0, "dsfs", "delhi")

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func getCalendar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	dbConn := ConnectDb()
	var calendar Calendar
	row := dbConn.QueryRowx("SELECT name, id, active, location, color FROM calendar where id =?", id)
	fmt.Println(row)
	err := row.StructScan(&calendar)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(Calendar(calendar))
}

func checkFunc(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(" ID is ", params["cid"])
	fmt.Println("CID is ", params["id"])
}

func main() {
	currentTime := time.Now()
	fmt.Println("Current date is ", currentTime)
	r := mux.NewRouter()
	r.HandleFunc("/{cid}/calendar/{id}", checkFunc).Methods("GET")
	r.HandleFunc("/api/calendar/{id}", getCalendar).Methods("GET")
	r.HandleFunc("/api/calendar", CreateCalendar).Methods("POST")
	r.HandleFunc("/api/calendar/{id}", DeleteCalendar).Methods("DELETE")
	r.HandleFunc("/api/calendar/{id}", PatchCalendar).Methods("PATCH")
	http.ListenAndServe(":8080", r)
}
