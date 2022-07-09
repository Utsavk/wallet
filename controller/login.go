package controller

import (
	"encoding/json"
	"wallet/logs"
	"wallet/service"
	"wallet/wcontext"

	"github.com/valyala/fasthttp"
)

func OnLoginRequest(ctx *wcontext.Context) ([]byte, int) {
	fctx := ctx.Fctx
	var loginService = &service.LoginService{}
	if fctx.IsPost() {
		var loginReqBody = service.LoginRequest{}
		if err := json.Unmarshal(ctx.Fctx.PostBody(), &loginReqBody); err != nil {
			logs.Print(err.Error())
			return []byte("login request body could not be parsed"), fasthttp.StatusBadRequest
		}
		message, authHeader, err := loginService.Login(loginReqBody)
		if err != nil {
			var respMsg []byte
			if message != "" {
				respMsg = []byte(message)
			} else {
				respMsg = []byte("bad request")
			}
			return respMsg, fasthttp.StatusBadRequest
		}
		fctx.Response.Header.Add("Authorization", authHeader)
		return []byte(message), fasthttp.StatusOK
	}
	return []byte("method not allowed"), fasthttp.StatusMethodNotAllowed
}
