package service

import (
	"wallet/errors"
	"wallet/repository"
)

type LoginRequest struct {
	Username string
	Password string
}

type LoginServiceInterface interface {
	repository.SessionRepoInterface
	Login(args LoginRequest) (string, string, *errors.Err)
}

type LoginService struct {
	sessionRepo repository.SessionRepoInterface
}

func (u *LoginService) Login(args LoginRequest) (string, string, *errors.Err) {
	return "logged in successfuly", "authHeaderValue", nil
}
