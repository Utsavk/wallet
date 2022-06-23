package controller

import (
	"wallet/service"

	"github.com/valyala/fasthttp"
)

func OnUserRequest(ctx *fasthttp.RequestCtx) {
	userService := &service.User{}
	if ctx.IsGet() {
		user := userService.GetUserDetailsByID("bd911361-e15f-443f-b71d-eea9cb5d5e9e")
		ctx.Response.SetBody([]byte(user.ID))
		return
	}
	ctx.Response.SetBody([]byte("method not allowed"))
}
