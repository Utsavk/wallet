package errors

import "fmt"

type Err struct {
	OriginalError error
	LogMessage    string
	HttpMessage   string
	HttpCode      int
}

func NewLogError(err error, logMessage string, httpMessage string, httpCode int) Err {
	return Err{
		OriginalError: err,
		LogMessage:    httpMessage,
		HttpMessage:   httpMessage,
		HttpCode:      httpCode,
	}
}

func (e *Err) Print() {
	fmt.Println(e.LogMessage)
}
