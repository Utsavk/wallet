package repository

import (
	"wallet/errors"
	"wallet/models"
)

type SessionRepoInterface interface {
	GetDBSessionByToken(token string) (*models.Session, *errors.Err)
	CreateDBSession(userID uint) (*models.Session, *errors.Err)
	UpdateDBSessionByID(id uint) (*models.Session, *errors.Err)
	DeleteDBSessionByID(id uint) *errors.Err
	GetDBSessionByID(id uint) (*models.Session, *errors.Err)
	GetDBSessionByUserID(userID uint) (*models.Session, *errors.Err)
}

type SessionRepo struct{}

func (s *SessionRepo) GetDBSessionByToken(token string) (*models.Session, *errors.Err) {
	return &models.Session{}, nil
}

func (s *SessionRepo) CreateDBSession(userID uint) (*models.Session, *errors.Err) {
	return &models.Session{}, nil
}

func (s *SessionRepo) UpdateDBSessionByID(id uint) (*models.Session, *errors.Err) {
	return &models.Session{}, nil
}

func (s *SessionRepo) DeleteDBSessionByID(id uint) *errors.Err {
	return nil
}

func (s *SessionRepo) GetDBSessionByID(id uint) (*models.Session, *errors.Err) {
	return &models.Session{}, nil
}

func (s *SessionRepo) GetDBSessionByUserID(userID uint) (*models.Session, *errors.Err) {
	return &models.Session{}, nil
}
