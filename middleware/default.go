package middleware

import (
	"wallet/context"
	"wallet/errors"
)

type mwFunc func(ctx *context.Ctx) *errors.Err

var authMw = AuthMw{}

var pathMiddlewares = map[string][]mwFunc{
	"/user": {
		authMw.VerifyAuth,
	},
	"/logout": {
		authMw.VerifyAuth,
	},
}

func Filter(ctx *context.Ctx) *errors.Err {
	filters := pathMiddlewares[ctx.Route]
	// traverse through various middleware layers
	for _, filterFn := range filters {
		filterFn(ctx)
	}
	return nil
}
