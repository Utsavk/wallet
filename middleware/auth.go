package middleware

import "github.com/valyala/fasthttp"

type AuthMwInterface interface {
	VerifyAuth(ctx *fasthttp.RequestCtx) error
}

type AuthMw struct{}

func (authmw *AuthMw) VerifyAuth(ctx *fasthttp.RequestCtx) error {
	return nil
}
