package utils

import (
	"fmt"
	"time"
)

func DateRange(t time.Time, i string) (from time.Time, to time.Time, err error) {
	now := t
	currentYear, currentMonth, currentDay := now.Date()
	currentLocation := now.Location()

	if i == "month" {
		from = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		to = from.AddDate(0, 1, -1)
	} else if i == "year" {
		from = time.Date(currentYear, 1, 1, 0, 0, 0, 0, currentLocation)
		to = from.AddDate(0, 12, -1)
	} else if i == "day" {
		from = time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)
		to = now
	} else {
		err = fmt.Errorf("Not valid '%s'.\nThe valid ones are: month, year and day", i)
	}

	to = time.Date(to.Year(), to.Month(), to.Day(), 23, 59, 59, 0, currentLocation)
	return
}
