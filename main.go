package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	holiday "github.com/najeira/jpholiday"
	"golang.org/x/xerrors"
)

type Response struct {
	Bizday int `json:"bizday"`
}

func isWeekend(t time.Time) bool {
	dayOfWeek := t.Weekday()
	if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
		return true
	}
	return false
}

func countBizDayInDays(days int) int {
	t := time.Now()

	//今日も日数に含む
	count := 0

	for i := 0; i < days; i++ {
		if isWeekend(t) || isHoliday(t) {
			t = t.Add(time.Hour * 24)
			continue
		}
		count++
		t = t.Add(time.Hour * 24)
	}
	log.Println("count is", count)
	return count
}

func isHoliday(t time.Time) bool {
	return holiday.Name(t) != ""
}

func handleBizDay(w http.ResponseWriter, r *http.Request) {
	days, _ := strconv.Atoi(r.FormValue("days"))
	bdays := countBizDayInDays(days)
	resp, err := json.Marshal(Response{bdays})

	if err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func handleIsWeekday(w http.ResponseWriter, r *http.Request) {

	t, err := time.Parse("2006-1-2", r.FormValue("date"))
	if err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
	}
	if isWeekend(t) {
		log.Println("That day is weekend")
		return
	}
	log.Println("That day is weekday")
}

func main() {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/is-weekday", handleIsWeekday)
		r.Post("/biz-day", handleBizDay)
	})

	log.Println("Listening...")
	http.ListenAndServe(":8002", r)
}
