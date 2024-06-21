package service

import (
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	// determmine our date
	days := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	daysLeft := time.Until(days)

	return int64(daysLeft.Hours()) / 24
}
