package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/kotatabe/api_date_conversion/handler/utils"
)

type ResponseBizDay struct {
	Bizday int `json:"bizday"`
}

type ResponseIsWeekend struct {
	Is_weekend bool `json:"isWeekend"`
}

func HandleBizDay(w http.ResponseWriter, r *http.Request) {
	days, err := strconv.Atoi(r.FormValue("days"))
	if err != nil {
		fmt.Printf("%+v\n", fmt.Errorf(": %w", err))
		http.Error(w, "400 Status Bad Request", http.StatusBadRequest)
		return
	}

	biz_days := utils.CountBizDayFromToday(days)

	resp, err := json.Marshal(ResponseBizDay{biz_days})
	if err != nil {
		fmt.Printf("%+v\n", fmt.Errorf(": %w", err))
		http.Error(w, "400 Status Bad Request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func HandleIsWeekend(w http.ResponseWriter, r *http.Request) {
	date, err := time.Parse("2006-1-2", r.FormValue("date"))
	if err != nil {
		fmt.Printf("%+v\n", fmt.Errorf(": %w", err))
		http.Error(w, "400 Status Bad Request", http.StatusBadRequest)
		return
	}

	d := utils.Date{T: date}
	is_weekend := d.IsWeekend()
	resp, err := json.Marshal(ResponseIsWeekend{is_weekend})
	if err != nil {
		fmt.Printf("%+v\n", fmt.Errorf(": %w", err))
		http.Error(w, "400 Status Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
