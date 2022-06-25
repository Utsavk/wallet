package middleware

import (
	"wallet/wcontext"

	"github.com/valyala/fasthttp"
)

type mwFunc func(ctx *fasthttp.RequestCtx) error

var authMw = AuthMw{}

var pathMiddlewares = map[string][]mwFunc{
	"user": {
		authMw.VerifyAuth,
	},
	"logout": {
		authMw.VerifyAuth,
	},
}

func Filter(ctx *wcontext.Context) error {
	filters := pathMiddlewares[ctx.Route]
	// traverse through various middleware layers
	for _, filterFn := range filters {
		filterFn(ctx.Fctx)
	}
	return nil
}
