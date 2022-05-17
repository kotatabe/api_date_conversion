package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"golang.org/x/xerrors"
	"github.com/kotatabe/api_date_conversion/handler/utils"
)

type Response struct {
	Bizday int `json:"bizday"`
}

func handleBizDay(w http.ResponseWriter, r *http.Request) {
	days, _ := strconv.Atoi(r.FormValue("days"))
	bdays := utils.countBizDayInDays(days)
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
	if utils.isWeekend(t) {
		log.Println("That day is weekend")
		return
	}
	log.Println("That day is weekday")
}
