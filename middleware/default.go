package middleware

import (
	"wallet/wcontext"

	"github.com/valyala/fasthttp"
)

type mw func(ctx *fasthttp.RequestCtx) error

var pathMiddlewares = map[string][]mw{
	"user": {
		VerifyAuth,
	},
	"logout": {
		VerifyAuth,
	},
}

func Filter(ctx *wcontext.Context) error {
	// filters, found := pathMiddlewares[ctx.URI().String()]
	// traverse through various middleware layers
	return nil
}
