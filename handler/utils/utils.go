package utils

import (
	"log"
	"time"

	holiday "github.com/najeira/jpholiday"
)

type Date struct {}


func (d *Date) IsWeekend(t time.Time) bool {
	dayOfWeek := t.Weekday()
	if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
		return true
	}
	return false
}

func (d *Date) isHoliday(t time.Time) bool {
	return holiday.Name(t) != ""
}

func CountBizDayFromToday(days int) int {
	t := time.Now()
	d := Date{}

	count := 0
	for i := 0; i < days; i++ {
		if !d.IsWeekend(t) && !d.isHoliday(t) {
			count++
		}
		t = t.Add(time.Hour * 24)
	}
	log.Println("count is", count)
	return count
}
