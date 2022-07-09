package server

import (
	"strings"
	"wallet/controller"
	"wallet/middleware"
	"wallet/wcontext"

	"github.com/valyala/fasthttp"
)

func requestHandler(fctx *fasthttp.RequestCtx) {
	var body = []byte("page not found")
	var status int = fasthttp.StatusNotFound
	ctx := &wcontext.Context{
		Fctx: fctx,
	}
	ctx.Route = string(fctx.RequestURI())
	middleware.Filter(ctx)
	if strings.HasPrefix(ctx.Route, "/user") {
		body, status = controller.OnUserRequest(ctx)
	} else if strings.HasPrefix(ctx.Route, "/login") {
		body, status = controller.OnLoginRequest(ctx)
	}
	sendResponse(fctx, status, body)
}
