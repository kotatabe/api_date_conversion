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

type Response struct {
	Bizday int `json:"bizday"`
}

func HandleBizDay(w http.ResponseWriter, r *http.Request) {
	days, err := strconv.Atoi(r.FormValue("days"))
	if err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
	}
	biz_days := utils.CountBizDayInDays(days)
	resp, err := json.Marshal(Response{biz_days})

	if err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func HandleIsWeekday(w http.ResponseWriter, r *http.Request) {

	t, err := time.Parse("2006-1-2", r.FormValue("date"))
	if err != nil {
		fmt.Printf("%+v\n", xerrors.Errorf(": %w", err))
	}
	if utils.IsWeekend(t) {
		log.Println("That day is weekend")
		return
	}
	log.Println("That day is weekday")
}
