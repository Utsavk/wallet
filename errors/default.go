package errors

import (
	"wallet/wcontext"
)

type Err struct {
	OriginalError error
	LogMessage    string
}

func New() Err {
	return Err{}
}

func NewError(err error, logMessage string, User *wcontext.UserInfo) Err {
	return Err{
		OriginalError: err,
		LogMessage:    logMessage,
	}
}
