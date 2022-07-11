package service

import (
	"wallet/context"
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
	CreateSession(ctx *context.Ctx, userID uint) (*Session, *errors.Err)
	UpdateSessionByID(ctx *context.Ctx, id uint) (*Session, *errors.Err)
	DeleteSessionByID(ctx *context.Ctx, id uint) *errors.Err
	GetSessionByID(ctx *context.Ctx, id uint) (*Session, *errors.Err)
	GetSessionByUserID(ctx *context.Ctx, userID uint) (*Session, *errors.Err)
}

type SessionService struct {
	sessionRepo repository.SessionRepoInterface
}

func (s *SessionService) CreateSession(ctx *context.Ctx, userID uint) (*Session, *errors.Err) {
	if userID == 0 {
		return nil, errors.NewError(nil, "invalid userid", nil)
	}
	_, err := s.sessionRepo.CreateDBSession(ctx, userID)
	if err != nil {
		return nil, err
	}
	session, err := s.sessionRepo.GetDBSessionByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &Session{
		ID:       session.ID,
		UUID:     session.UUID,
		Token:    session.Token,
		UserID:   session.UserID,
		ExpiryAt: session.ExpiryAt,
	}, nil
}

func (s *SessionService) UpdateSessionByID(ctx *context.Ctx, id uint) (*Session, *errors.Err) {
	return &Session{}, nil
}

func (s *SessionService) DeleteSessionByID(ctx *context.Ctx, id uint) *errors.Err {
	return nil
}

func (s *SessionService) GetSessionByID(ctx *context.Ctx, id uint) (*Session, *errors.Err) {
	return &Session{}, nil
}

func (s *SessionService) GetSessionByUserID(ctx *context.Ctx, userID string) (*Session, *errors.Err) {
	return &Session{}, nil
}

func (s *SessionService) GetSessionByToken(ctx *context.Ctx, token string) (*Session, *errors.Err) {
	return &Session{}, nil
}
