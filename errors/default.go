package errors

import (
	"wallet/context"
)

type Err struct {
	OriginalError error
	LogMessage    string
	UserID        uint
	MessageCode   string // prepare an enum for it
}

func New() *Err {
	return &Err{}
}

func NewError(err error, logMessage string, User *context.UserInfo) *Err {
	errObj := &Err{
		OriginalError: err,
		LogMessage:    logMessage,
	}
	if User != nil {
		errObj.UserID = User.ID
	}
	return errObj
}
