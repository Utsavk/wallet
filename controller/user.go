package controller

import (
	"encoding/json"
	"strconv"
	"wallet/logs"
	"wallet/service"
	"wallet/wcontext"

	"github.com/valyala/fasthttp"
)

func OnUserRequest(ctx *wcontext.Context) ([]byte, int) {
	userService := &service.UserService{}
	fctx := ctx.Fctx
	if fctx.IsGet() {
		id, err := strconv.Atoi(string(fctx.QueryArgs().Peek("id")))
		if err != nil {
			logs.Print(err.Error())
			return []byte("invalid query args"), fasthttp.StatusBadRequest
		}
		user := userService.GetUserDetailsByID(id)
		userBytes, err := json.Marshal(user)
		if err != nil {
			logs.Print(err.Error())
			return []byte("server error"), fasthttp.StatusInternalServerError
		}
		return userBytes, fasthttp.StatusOK
	}
	if fctx.IsPost() {
		var userArgs = service.NewUserArgs{}
		if err := json.Unmarshal(ctx.Fctx.PostBody(), &userArgs); err != nil {
			logs.Print(err.Error())
			return []byte("create user request body could not be parsed"), fasthttp.StatusBadRequest
		}
		user := userService.CreateUser(userArgs)
		if user == nil {
			return []byte("user could not be created"), fasthttp.StatusBadRequest
		}
		userBytes, err := json.Marshal(user)
		if err != nil {
			logs.Print(err.Error())
			return []byte("server error"), fasthttp.StatusInternalServerError
		}
		return userBytes, fasthttp.StatusCreated
	}
	return []byte("method not allowed"), fasthttp.StatusMethodNotAllowed
}
