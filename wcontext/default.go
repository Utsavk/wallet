package wcontext

import "github.com/valyala/fasthttp"

type UserInfo struct {
	ID uint
}

type Context struct {
	Fctx  *fasthttp.RequestCtx
	Route string
	User  UserInfo
}
