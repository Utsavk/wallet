package middleware

import (
	"fmt"
	"time"
	"wallet/context"
	err "wallet/errors"
	"wallet/repository"

	"github.com/valyala/fasthttp"
)

type AuthMwInterface interface {
	repository.SessionRepoInterface
	VerifyAuth(ctx *context.Ctx) *err.Err
}

type AuthMw struct {
	sessionRepo repository.SessionRepoInterface
}

func (authmw *AuthMw) VerifyAuth(ctx *context.Ctx) *err.Err {
	token := getToken(ctx.Fctx)
	if token == "" {
		return err.NewError(nil, "auth token is not provided", nil)
	}
	session, err1 := authmw.sessionRepo.GetDBSessionByToken(ctx, token)
	expiryDate, _ := time.Parse("2006-01-02 15:04:05", session.ExpiryAt)
	if time.Now().After(expiryDate) {
		return err.NewError(nil, fmt.Sprintf("auth token %s is expired", token), &context.UserInfo{ID: session.UserID})
	}
	return err1
}

func getToken(ctx *fasthttp.RequestCtx) string {
	return string(ctx.Request.Header.Peek("Authorization"))
}
