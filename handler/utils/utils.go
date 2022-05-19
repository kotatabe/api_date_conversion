package utils

import (
	"log"
	"time"

	holiday "github.com/najeira/jpholiday"
)

type Date struct {
	T time.Time
}

func (d *Date) IsWeekend() bool {
	dayOfWeek := d.T.Weekday()
	if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
		return true
	}
	return false
}

func (d *Date) IsHoliday() bool {
	return holiday.Name(d.T) != ""
}

func (d *Date) AddOneDay() {
	d.T = d.T.Add(time.Hour * 24)
}

func CountBizDayFromToday(days int) int {
	d := Date{T: time.Now()}

	count := 0
	for i := 0; i < days; i++ {
		if !d.IsWeekend() && !d.IsHoliday() {
			count++
		}
		d.AddOneDay()
	}
	log.Println("count is", count)
	return count
}
