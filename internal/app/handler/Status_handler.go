package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// create an interface for dependency inject here
type Service interface {
	DaysLeft() int64
}

type StatusHandler struct {
	s Service
}

func NewStatusHandler(s Service) *StatusHandler {
	return &StatusHandler{
		s: s,
	}
}

func (sh *StatusHandler) GetStatus(ctx echo.Context) error {
	daysLeft := sh.s.DaysLeft()

	displayedMessage := fmt.Sprintf("Days number: %d", daysLeft)

	// send our user message
	err := ctx.String(http.StatusOK, displayedMessage)
	if err != nil {
		return err
	}

	return nil
}
