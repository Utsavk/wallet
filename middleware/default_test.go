package middleware

import (
	"testing"
	"wallet/middleware/mocks"
	"wallet/wcontext"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

func getMockContext(uri string, route string) *wcontext.Context {
	fctx := &fasthttp.RequestCtx{
		Request: fasthttp.Request{},
	}
	fctx.Request.SetRequestURI(uri)
	return &wcontext.Context{
		Fctx:  fctx,
		Route: route,
	}
}

func TestFilter(t *testing.T) {
	route := "user"
	ctx := getMockContext("/v1/"+route, route)
	authMw := mocks.AuthMwInterface{}
	var originalRouteFnList = pathMiddlewares[route]
	pathMiddlewares[route] = []mwFunc{
		authMw.VerifyAuth,
	}
	authMw.On("VerifyAuth", mock.Anything).Return(nil)
	gotErr := Filter(ctx)
	assert.Equal(t, nil, gotErr)
	pathMiddlewares[route] = originalRouteFnList
	authMw.AssertNumberOfCalls(t, "VerifyAuth", 1)
}
