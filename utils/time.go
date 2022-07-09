package utils

import "time"

const DB_TIME_FORMAT = "2006/01/02 15:04:05"

type ClockInterface interface {
	Now() time.Time
	GetCurrentTime() string
}

type Clock struct{}

var ClockObj ClockInterface = &Clock{}

func (c *Clock) Now() time.Time {
	return time.Now()
}

func (c *Clock) GetCurrentTime() string {
	return c.Now().Format(DB_TIME_FORMAT)
}
