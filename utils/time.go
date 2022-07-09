package utils

import "time"

const DB_TIME_FORMAT = "2006/01/02 15:04:05"

func GetCurrentTime() string {
	return time.Now().Format(DB_TIME_FORMAT)
}
