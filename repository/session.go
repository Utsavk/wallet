package repository

import (
	"wallet/errors"
	"wallet/models"
)

type SessionRepoInterface interface {
	GetSessionByToken(token string) (*models.Session, *errors.Err)
}

type SessionRepo struct {
}

func (s *SessionRepo) GetSessionByToken(token string) (*models.Session, *errors.Err) {
	return &models.Session{}, nil
}
