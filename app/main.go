package main

import (
	"github.com/hoangvantuan/go-base/container"
	"github.com/hoangvantuan/go-base/handler/hhttp"
	"github.com/hoangvantuan/go-base/infra"
	"github.com/hoangvantuan/go-base/usercase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {
	infra.BootConfig()
	infra.BootMysql()
	ctn := container.Boot()

	defer func() {
		err := infra.DB.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"*"},
	}))

	// wrap echo context to customize context
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			return h(&hhttp.Context{
				Context: ctx,
			})
		}
	})

	// middlware
	e.Use(hhttp.BindAuthMiddleware())
	// validator
	// e.Validator = hhttp.NewDefaultValidator()

	userUsecase := ctn.Get("user-usecase").(usercase.UserUsecase)

	hhttp.BindHttpUserHandler(e, userUsecase)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))

}
