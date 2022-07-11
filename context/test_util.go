package context

import (
	"math/rand"

	"github.com/valyala/fasthttp"
)

func GetMockContext(fctx *fasthttp.RequestCtx, user *UserInfo) *Ctx {
	if user == nil {
		user = &UserInfo{
			uint(rand.Uint32()),
		}
	}
	if fctx == nil {
		// create mock context of fasthttp
	}
	return &Ctx{User: user, Fctx: fctx}
}
