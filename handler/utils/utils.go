package utils

import (
	"log"
	"time"

	holiday "github.com/najeira/jpholiday"
)

func IsWeekend(t time.Time) bool {
	dayOfWeek := t.Weekday()
	if dayOfWeek == time.Saturday || dayOfWeek == time.Sunday {
		return true
	}
	return false
}

func isHoliday(t time.Time) bool {
	return holiday.Name(t) != ""
}

func CountBizDayInDays(days int) int {
	t := time.Now()

	//今日も日数に含む
	count := 0

	for i := 0; i < days; i++ {
		if IsWeekend(t) || isHoliday(t) {
			t = t.Add(time.Hour * 24)
			continue
		}
		count++
		t = t.Add(time.Hour * 24)
	}
	log.Println("count is", count)
	return count
}
