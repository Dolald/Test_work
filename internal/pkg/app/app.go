package app

import (
	"log"
	handler "test/internal/app/handler"
	service "test/internal/app/service"
	"test/internal/middleware"

	"github.com/labstack/echo/v4"
)

type App struct {
	sh   *handler.StatusHandler
	s    *service.Service
	echo *echo.Echo
}

func NewApp() (*App, error) {
	app := &App{}

	app.s = service.NewService()
	app.sh = handler.NewStatusHandler(app.s)

	echo := echo.New()
	// check user status
	echo.Use(middleware.CheckRole)

	err := echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	app.echo.GET("/status", app.sh.GetStatus)

	return app, nil
}

func (a *App) Run() error {
	// open our working port
	if err := a.echo.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	return nil
}
