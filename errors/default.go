package errors

import (
	"wallet/wcontext"
)

type Err struct {
	OriginalError error
	LogMessage    string
	UserID        string
}

func New() *Err {
	return &Err{}
}

func NewError(err error, logMessage string, User *wcontext.UserInfo) *Err {
	errObj := &Err{
		OriginalError: err,
		LogMessage:    logMessage,
	}
	if User != nil {
		errObj.UserID = User.ID
	}
	return errObj
}
