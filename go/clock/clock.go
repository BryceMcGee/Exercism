package clock

import (
	"fmt"
)

type Clock struct {
	hours   int
	minutes int
}

func adjustMin(hour, min int) (int, int) {
	for min > 59 {
		min = min - 60
		hour++
	}
	for min < 0 {
		hour--
		min = min + 60
	}
	return hour, min
}

func adjustHour(hour int) int {
	for hour > 23 {
		hour = hour - 24
	}

	for hour < 0 {
		hour = hour + 24
	}
	return hour
}

func New(h, m int) Clock {

	if m > 59 || m < 0 {
		h, m = adjustMin(h, m)
	}

	if h > 23 || h < 0 {
		h = adjustHour(h)
	}

	c := Clock{h, m}
	return c
}

func (c Clock) Add(m int) Clock {
	c.minutes += m

	if c.minutes > 59 {
		c.hours, c.minutes = adjustMin(c.hours, c.minutes)
	}

	if c.hours >= 24 {
		c.hours = adjustHour(c.hours)
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.minutes -= m

	if c.minutes < 0 {
		c.hours, c.minutes = adjustMin(c.hours, c.minutes)
	}

	if c.hours < 0 {
		c.hours = adjustHour(c.hours)
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
