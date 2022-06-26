package logs

import (
	"fmt"
	"time"
	"wallet/wcontext"
)

func Print(message string) {
	PrintDetail(message, nil)
}

func PrintDetail(message string, user *wcontext.UserInfo) {
	var logMessage = time.Now().Format("2006/01/02 15:04:05")
	if user != nil && user.ID != "" {
		logMessage += fmt.Sprintf("UserID: %s | ", user.ID)
	}
	logMessage += message + "\n"
	fmt.Print(logMessage)
}
