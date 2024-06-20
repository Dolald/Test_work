package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	handle()

	echo := echo.New()

	echo.Use(middleware)

	err := echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func sendMessage(e echo.Context) error {
	return e.String(http.StatusOK, "red button user detected")
}

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// get user role
		userRole := ctx.Request().Header.Get("User-role")

		if userRole == "admin" {
			log.Print("red button user detected")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
