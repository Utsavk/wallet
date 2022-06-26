package middleware

import (
	"fmt"
	"testing"
	"wallet/errors"
	"wallet/middleware/mocks"
	"wallet/wcontext"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

func getMWTestMockContext(uri string, route string) *wcontext.Context {
	fctx := &fasthttp.RequestCtx{
		Request: fasthttp.Request{},
	}
	fctx.Request.SetRequestURI(uri)
	return &wcontext.Context{
		Fctx:  fctx,
		Route: route,
	}
}

func TestFilter_ForVerifyAuth(t *testing.T) {
	type args struct {
		route string
	}
	type filterTest struct {
		name                string
		args                args
		wantVerifyAuthCount int
		err                 *errors.Err
	}

	testData := []filterTest{
		{
			name: "user route happy flow test",
			args: args{
				route: "/user",
			},
			wantVerifyAuthCount: 1,
			err:                 nil,
		},
		{
			name: "logout route happy flow test",
			args: args{
				route: "/logout",
			},
			wantVerifyAuthCount: 1,
			err:                 nil,
		},
	}

	for _, tdt := range testData {
		fmt.Println(tdt.name)
		route := tdt.args.route
		ctx := getMWTestMockContext("/v1"+route, route)
		authMw := mocks.AuthMwInterface{}
		var originalRouteFnList = pathMiddlewares[route]
		pathMiddlewares[route] = []mwFunc{
			authMw.VerifyAuth,
		}
		authMw.On("VerifyAuth", mock.Anything).Return(nil)
		gotErr := Filter(ctx)
		assert.Equal(t, tdt.err, gotErr)
		pathMiddlewares[route] = originalRouteFnList
		authMw.AssertNumberOfCalls(t, "VerifyAuth", tdt.wantVerifyAuthCount)
	}
}

func TestFilter_ForNoMW(t *testing.T) {
	type args struct {
		route string
	}
	type filterTest struct {
		name                string
		args                args
		wantVerifyAuthCount int
		err                 *errors.Err
	}

	testData := []filterTest{
		{
			name: "login route happy flow test",
			args: args{
				route: "/login",
			},
			wantVerifyAuthCount: 0,
			err:                 nil,
		},
	}

	for _, tdt := range testData {
		fmt.Println(tdt.name)
		route := tdt.args.route
		ctx := getMWTestMockContext("/v1"+route, route)
		authMw := mocks.AuthMwInterface{}
		authMw.On("VerifyAuth", mock.Anything).Return(nil)
		gotErr := Filter(ctx)
		assert.Equal(t, tdt.err, gotErr)
		authMw.AssertNumberOfCalls(t, "VerifyAuth", tdt.wantVerifyAuthCount)
	}
}
