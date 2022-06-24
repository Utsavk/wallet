package server

import "github.com/valyala/fasthttp"

func sendResponse(fctx *fasthttp.RequestCtx, code int, body []byte) {
	if body != nil {
		fctx.Response.SetBody(body)
	}
	if code != 0 {
		fctx.SetStatusCode(code)
	}
}
