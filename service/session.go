package service

import (
	"wallet/errors"
	"wallet/repository"
)

type Session struct {
	ID         uint
	UUID       string
	Token      string
	UserID     uint
	LastActive string
	ExpiryAt   string
}

type SessionServiceInterface interface {
	repository.SessionRepoInterface
	CreateSession(userID uint) (*Session, *errors.Err)
	UpdateSessionByID(id uint) (*Session, *errors.Err)
	DeleteSessionByID(id uint) *errors.Err
	GetSessionByID(id uint) (*Session, *errors.Err)
	GetSessionByUserID(userID uint) (*Session, *errors.Err)
}

type SessionService struct {
	sessionRepo repository.SessionRepoInterface
}

func (s *SessionService) CreateSession(userID uint) (*Session, *errors.Err) {
	if userID == 0 {
		return nil, errors.NewError(nil, "invalid userid", nil)
	}
	session, err := s.sessionRepo.CreateDBSession(userID)
	if err != nil {
		return nil, err
	}
	return &Session{
		ID:     session.ID,
		UUID:   session.UUID,
		Token:  session.Token,
		UserID: session.UserID,
	}, nil
}

func (s *SessionService) UpdateSessionByID(id uint) (*Session, *errors.Err) {
	return &Session{}, nil
}

func (s *SessionService) DeleteSessionByID(id uint) *errors.Err {
	return nil
}

func (s *SessionService) GetSessionByID(id uint) (*Session, *errors.Err) {
	return &Session{}, nil
}

func (s *SessionService) GetSessionByUserID(userID string) (*Session, *errors.Err) {
	return &Session{}, nil
}

func (s *SessionService) GetSessionByToken(token string) (*Session, *errors.Err) {
	return &Session{}, nil
}
