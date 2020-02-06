package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type CalendarRequest struct {
	Location    string    `json:"location"`
	Update_dt   time.Time `json:"update_dt"`
	Id          string    `json:"id"`
	Creation_dt time.Time `json:"creation_dt"`
	Name        *string   `json:"name"`
	Active      *bool     `json:"active"`
	Color       *int      `json:"color"`
	Overlap     *bool     `json:"overlap"`
	Attributes  *string   `json:"attributes"`
	Image_Url   *string   `json:"image_url"`
}

type EventRequest struct {
	Start_dt      *time.Time `json:"start_dt"`
	End_dt        *time.Time `json:"end_dt"`
	All_day       *bool      `json:"all_day"`
	Title         *string    `json:"title"`
	Who           *string    `json:"who"`
	Location      *string    `json:"location"`
	Notes         *string    `json:"notes"`
	Id            string     `json:"id"`
	SubCalendarId string     `json:"subcalendarid"`
	CalendarId    string     `json:"calendarid"`
	Version       string     `json:"version"`
	Creation_dt   time.Time  `json:"creation_dt"`
	Update_dt     time.Time  `json:"update_dt"`
}

type ErrorStruct struct {
	isError     bool
	ErrorString string
}

// insert into calendar(id, name, active, color, overlap, attributes, location) values ('123', 'sailesh', 1, 4, 0, "dsfs", "delhi")

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	currentTime := time.Now()
	fmt.Println("Current date is ", currentTime.Format("2006-01-02 15:04:05"))

	r := mux.NewRouter()
	r.HandleFunc("/api/calendar/{id}", GetCalendar).Methods("GET")
	r.HandleFunc("/api/calendar", CreateCalendar).Methods("POST")
	r.HandleFunc("/api/calendar/{id}", DeleteCalendar).Methods("DELETE")
	r.HandleFunc("/api/calendar/{id}", PatchCalendar).Methods("PATCH")
	r.HandleFunc("/{calendarId}/events", CreateEvent).Methods("POST")
	r.HandleFunc("/{calendarId}/events/{eventId}", GetEvent).Methods("GET")
	r.HandleFunc("/{calendarId}/events", GetEventRange).Methods("GET")
	http.ListenAndServe(":8080", r)
}
