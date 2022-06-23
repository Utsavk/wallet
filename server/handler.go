package server

import (
	"wallet/controller"

	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	uri := string(ctx.RequestURI())
	if uri == "/user" {
		controller.OnUserRequest(ctx)
		return
	}
	ctx.Response.SetBody([]byte("page not found"))
}
