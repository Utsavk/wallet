package middleware

import (
	"wallet/errors"
	"wallet/wcontext"

	"github.com/valyala/fasthttp"
)

type mwFunc func(ctx *fasthttp.RequestCtx) *errors.Err

var authMw = AuthMw{}

var pathMiddlewares = map[string][]mwFunc{
	"/user": {
		authMw.VerifyAuth,
	},
	"/logout": {
		authMw.VerifyAuth,
	},
}

func Filter(ctx *wcontext.Context) *errors.Err {
	filters := pathMiddlewares[ctx.Route]
	// traverse through various middleware layers
	for _, filterFn := range filters {
		filterFn(ctx.Fctx)
	}
	return nil
}
