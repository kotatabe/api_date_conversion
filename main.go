package main

import (
	"log"
	"net/http"
	"time"
)

func isWeekend(t time.Time) bool {
	dayOfWeek := t.Weekday()
	if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
		return true
	}
	return false
}

func handleSample(w http.ResponseWriter, r *http.Request) {

	t, _ := time.Parse("2006-01-02", r.FormValue("date"))

	if isWeekend(t) {
		log.Println("That day is weekend")
		return
	}
	log.Println("That day is weekday")
}

func main() {

	http.HandleFunc("/", handleSample)

	http.ListenAndServe(":8002", nil)
}
