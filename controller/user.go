package controller

import (
	"encoding/json"
	"fmt"
	"strconv"
	"wallet/context"
	"wallet/logs"
	"wallet/service"

	"github.com/valyala/fasthttp"
)

func OnUserRequest(ctx *context.Ctx) ([]byte, int) {
	userService := &service.UserService{}
	fctx := ctx.Fctx
	if fctx.IsGet() {
		id, err := strconv.Atoi(string(fctx.QueryArgs().Peek("id")))
		if err != nil {
			logs.Print(err.Error())
			return []byte("invalid query args"), fasthttp.StatusBadRequest
		}
		user, err1 := userService.GetUserDetailsByID(ctx, id)
		if err1 != nil { // check for enum
			return []byte(err1.LogMessage), fasthttp.StatusBadRequest
		}
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
		userId, err1 := userService.CreateUser(ctx, userArgs)
		if err1 != nil {
			return []byte("user could not be created"), fasthttp.StatusBadRequest
		}
		return []byte(fmt.Sprintf("new user with id %d is created", userId)), fasthttp.StatusCreated
	}
	return []byte("method not allowed"), fasthttp.StatusMethodNotAllowed
}
