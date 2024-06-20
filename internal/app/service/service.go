package service

import (
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) daysLeft() int {
	// determmine our date
	days := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	daysLeft := time.Until(days)

	return int(daysLeft.Hours()) / 24
}
