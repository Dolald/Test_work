package middleware

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

func CheckRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// get user role
		userRole := ctx.Request().Header.Get("User-role")

		if strings.EqualFold(userRole, "admin") {
			log.Print("red button user detected")
		}
		// define next middleware
		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
