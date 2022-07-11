package logs

import (
	"fmt"
	"time"
	"wallet/context"
)

func Print(message interface{}) {
	PrintDetail(message, nil)
}

func PrintDetail(message interface{}, user *context.UserInfo) {
	var logMessage = time.Now().Format("2006/01/02 15:04:05") + " "
	if user != nil && user.ID != 0 {
		logMessage += fmt.Sprintf("UserID: %d | ", user.ID)
	}
	logMessage += fmt.Sprintf("%v\n", message)
	fmt.Print(logMessage)
}
