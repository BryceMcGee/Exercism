package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	t, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	t, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	h := t.Hour()
	return h >= 12 && h <= 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	t := Schedule(date)
	wd := t.Weekday()
	m := t.Month()
	d := t.Day()
	y := t.Year()
	hr, min, _ := t.Clock()

	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", wd, m, d, y, hr, min)
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	y, _, _ := time.Now().Date()
	date := fmt.Sprintf("%v-09-15", y)
	t, _ := time.Parse("2006-01-02", date)
	return t
}
