package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Service interface {
}

type StatusHandler struct {
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}

func (sh *StatusHandler) GetStatus(ctx echo.Context, daysLeft int) error {
	displayedMessage := fmt.Sprintf("Days number: %d", daysLeft)

	// send our user message
	if err := ctx.String(200, displayedMessage); err != nil {
		return err
	}

	return nil
}
