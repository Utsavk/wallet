package server

import (
	"wallet/controller"
	"wallet/middleware"
	"wallet/wcontext"

	"github.com/valyala/fasthttp"
)

func requestHandler(fctx *fasthttp.RequestCtx) {
	ctx := &wcontext.Context{
		Fctx: fctx,
	}
	ctx.Route = string(fctx.RequestURI())
	middleware.Filter(ctx)
	if ctx.Route == "/user" {
		controller.OnUserRequest(ctx)
		return
	}
	sendResponse(fctx, fasthttp.StatusNotFound, []byte("page not found"))
}
