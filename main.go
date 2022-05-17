package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kotatabe/api_date_conversion/handler"
)

func main() {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/is-weekday", handler.HandleIsWeekday)
		r.Post("/biz-day", handler.HandleBizDay)
	})

	log.Println("Listening...")
	http.ListenAndServe(":8002", r)
}
