package repository

import (
	"wallet/context"
	"wallet/errors"
	"wallet/models"
)

type SessionRepoInterface interface {
	GetDBSessionByToken(ctx *context.Ctx, token string) (*models.Session, *errors.Err)
	CreateDBSession(ctx *context.Ctx, userID uint) (int64, *errors.Err)
	UpdateDBSessionByID(ctx *context.Ctx, id uint) *errors.Err
	DeleteDBSessionByID(ctx *context.Ctx, id uint) *errors.Err
	GetDBSessionByID(ctx *context.Ctx, id uint) (*models.Session, *errors.Err)
	GetDBSessionByUserID(ctx *context.Ctx, userID uint) (*models.Session, *errors.Err)
}

type SessionRepo struct{}

func (s *SessionRepo) GetDBSessionByToken(ctx *context.Ctx, token string) (*models.Session, *errors.Err) {
	return &models.Session{}, nil
}

func (s *SessionRepo) CreateDBSession(ctx *context.Ctx, userID uint) (int64, *errors.Err) {
	return 0, nil
}

func (s *SessionRepo) UpdateDBSessionByID(ctx *context.Ctx, id uint) *errors.Err {
	return nil
}

func (s *SessionRepo) DeleteDBSessionByID(ctx *context.Ctx, id uint) *errors.Err {
	return nil
}

func (s *SessionRepo) GetDBSessionByID(ctx *context.Ctx, id uint) (*models.Session, *errors.Err) {
	return &models.Session{}, nil
}

func (s *SessionRepo) GetDBSessionByUserID(ctx *context.Ctx, userID uint) (*models.Session, *errors.Err) {
	return &models.Session{}, nil
}
