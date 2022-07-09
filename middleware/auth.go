package middleware

import (
	"fmt"
	"time"
	err "wallet/errors"
	"wallet/repository"
	"wallet/wcontext"

	"github.com/valyala/fasthttp"
)

type AuthMwInterface interface {
	repository.SessionRepoInterface
	VerifyAuth(ctx *fasthttp.RequestCtx) *err.Err
}

type AuthMw struct {
	sessionRepo repository.SessionRepoInterface
}

func (authmw *AuthMw) VerifyAuth(ctx *fasthttp.RequestCtx) *err.Err {
	token := getToken(ctx)
	if token == "" {
		return err.NewError(nil, "auth token is not provided", nil)
	}
	session, err1 := authmw.sessionRepo.GetDBSessionByToken(token)
	expiryDate, _ := time.Parse("2006-01-02 15:04:05", session.ExpiryAt)
	if time.Now().After(expiryDate) {
		return err.NewError(nil, fmt.Sprintf("auth token %s is expired", token), &wcontext.UserInfo{ID: session.UserID})
	}
	return err1
}

func getToken(ctx *fasthttp.RequestCtx) string {
	return string(ctx.Request.Header.Peek("Authorization"))
}
