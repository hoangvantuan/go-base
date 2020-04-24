package hhttp

import (
	"github.com/labstack/echo"
)

func BindAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return do(func(ctx *Context) error {
			token := ctx.Request().Header.Get("Authorization")
			err := checkToken(ctx, token)
			if err != nil {
				return e(ctx, err)
			}

			return next(ctx)
		})
	}
}

func checkToken(ctx *Context, token string) error {
	if token == "" {
		return newTokenNotFound()
	}

	return nil
}