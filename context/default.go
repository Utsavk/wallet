package context

import "github.com/valyala/fasthttp"

type UserInfo struct {
	ID uint
}

type Ctx struct {
	Fctx  *fasthttp.RequestCtx
	Route string
	User  *UserInfo
}
