package hhttp

import (
	"net/http"

	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
}

type callFunc func(c *Context) error

func do(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*Context))
	}
}

type response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func e(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusOK, &response{
		Success: false,
		Error:   err.Error(),
	})

	return nil
}

func s(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, &response{
		Success: true,
		Data:    data,
	})
}
