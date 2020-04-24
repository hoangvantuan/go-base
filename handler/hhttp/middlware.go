package hhttp

import (
	"strings"

	"github.com/labstack/echo"
)

// no need auth path
var excludes = []string{"/"}

// BindAuthMiddleware bind authentication middlware use JWT token
func BindAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return do(func(ctx *Context) error {
			path := ctx.Path()
			for _, e := range excludes {
				if e == path {
					return next(ctx)
				}
			}

			bearerToken := ctx.Request().Header.Get("Authorization")
			token := strings.TrimPrefix(bearerToken, "Bearer ")
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
	// TODO need check more condition
	return nil
}
