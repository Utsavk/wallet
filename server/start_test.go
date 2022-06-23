package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func getMockCtx(route string) *fasthttp.RequestCtx {
	ctx := &fasthttp.RequestCtx{
		Request: fasthttp.Request{},
	}
	ctx.Request.SetRequestURI("/" + route)
	return ctx
}

func TestRequestHandler(t *testing.T) {
	type args struct {
		mockCtx *fasthttp.RequestCtx
	}
	type row struct {
		args        args
		wantErr     bool
		wantResBody []byte
	}
	rows := []row{
		{
			args: args{
				mockCtx: getMockCtx("user"),
			},
			wantErr:     false,
			wantResBody: []byte("Hello world"),
		},
		{
			args: args{
				mockCtx: getMockCtx("random"),
			},
			wantErr:     false,
			wantResBody: []byte("page not found"),
		},
	}

	for _, tt := range rows {
		requestHandler(tt.args.mockCtx)
		assert.Equal(t, tt.wantResBody, tt.args.mockCtx.Response.Body())
	}
}
