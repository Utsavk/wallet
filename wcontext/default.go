package wcontext

import "github.com/valyala/fasthttp"

type Context struct {
	Fctx  *fasthttp.RequestCtx
	Route string
}
