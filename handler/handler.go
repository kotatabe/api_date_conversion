package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kotatabe/api_date_conversion/handler/utils"
	"golang.org/x/xerrors"
)

type ResponseBizDay struct {
	Bizday int `json:"bizday"`
}

type ResponseIsWeekend struct {
	is_weekend bool `json:"isWeekend"`
}

func HandleBizDay(w http.ResponseWriter, r *http.Request) {
	days, err := strconv.Atoi(r.FormValue("days"))
	if err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
		http.Error(w, fmt.Sprintf("...: %w", err), http.StatusInternalServerError)
		return
	}

	biz_days := utils.CountBizDayFromToday(days)

	resp, err := json.Marshal(ResponseBizDay{biz_days})
	if err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
		http.Error(w, fmt.Sprintf("...: %w", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func HandleIsWeekend(w http.ResponseWriter, r *http.Request) {
	d := utils.Date{}
	t, err := time.Parse("2006-1-2", r.FormValue("date"))
	if err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
		http.Error(w, fmt.Sprintf("...: %w", err), http.StatusInternalServerError)
		return
	}

	is_weekend := d.IsWeekend(t)
	log.Println(is_weekend)
	resp, err := json.Marshal(ResponseIsWeekend{is_weekend})
	if err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
		http.Error(w, fmt.Sprintf("...: %w", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
