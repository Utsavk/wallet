package controller

import (
	"wallet/service"
	"wallet/wcontext"
)

func OnUserRequest(ctx *wcontext.Context) {
	userService := &service.User{}
	fctx := ctx.Fctx
	if fctx.IsGet() {
		user := userService.GetUserDetailsByID("bd911361-e15f-443f-b71d-eea9cb5d5e9e")
		fctx.Response.SetBody([]byte(user.ID))
		return
	}
	fctx.Response.SetBody([]byte("method not allowed"))
}
