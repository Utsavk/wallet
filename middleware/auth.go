package middleware

import (
	"wallet/errors"
	"wallet/logs"
	"wallet/repository"

	"github.com/valyala/fasthttp"
)

type AuthMwInterface interface {
	repository.SessionRepoInterface
	VerifyAuth(ctx *fasthttp.RequestCtx) *errors.Err
}

type AuthMw struct {
	sessionRepo repository.SessionRepoInterface
}

func (authmw *AuthMw) VerifyAuth(ctx *fasthttp.RequestCtx) *errors.Err {
	session, err := authmw.sessionRepo.GetSessionByToken(getToken(ctx))
	logs.Print(session.Token)
	return err
}

func getToken(ctx *fasthttp.RequestCtx) string {
	return string(ctx.Request.Header.Peek("Authorization"))
}
