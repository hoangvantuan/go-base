package hhttp

import (
	"github.com/hoangvantuan/go-base/helper"
	"github.com/hoangvantuan/go-base/logger"
	"github.com/hoangvantuan/go-base/usercase"
	"github.com/labstack/echo"
)

// BindHttpUserHandler bind user usecase for handler
func BindHttpUserHandler(ctx *echo.Echo, userUsecase usercase.UserUsecase) {
	uh := userHandler{
		userUsercase: userUsecase,
	}

	ctx.GET("/", do(uh.root))
	ctx.GET("/v1/users", do(uh.fetchAll))
}

type userHandler struct {
	userUsercase usercase.UserUsecase
}

func (uh *userHandler) fetchAll(ctx *Context) error {
	users, err := uh.userUsercase.FindAll()
	if err != nil {
		logger.Warn(helper.Pf("failed to find all user %s", err))
		return e(ctx, newSystemError())
	}

	if users == nil {
		return e(ctx, newDataNotFound())
	}

	return s(ctx, users)
}

func (uh *userHandler) root(ctx *Context) error {
	return s(ctx, "miichisoft quiz api is working :)")
}
