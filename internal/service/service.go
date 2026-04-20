package service

import (
	"fmt"
	"time"

	"github.com/Reteger/shiki/internal/models"
	"github.com/Reteger/shiki/internal/repository"
)

type Service interface {
	GetOngoings(days int) (*models.OngoingResponse, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetOngoings(days int) (*models.OngoingResponse, error) {
	if days < 1 || days > 7 {
		return nil, fmt.Errorf("days must be 1-7")
	}

	today := time.Now().Weekday() // 0..6
	// Преобразуем к int, считаем с модулем 7, потом обратно к Weekday
	targetIdx := (int(today) + days) % 7
	target := time.Weekday(targetIdx)

	dayNames := map[time.Weekday]string{
		time.Monday:    "Понедельник",
		time.Tuesday:   "Вторник",
		time.Wednesday: "Среда",
		time.Thursday:  "Четверг",
		time.Friday:    "Пятница",
		time.Saturday:  "Суббота",
		time.Sunday:    "Воскресенье",
	}

	targetName, ok := dayNames[target]
	if !ok {
		return nil, fmt.Errorf("unknown weekday: %v", target)
	}

	titles, err := s.repo.GetForDay(targetName)
	if err != nil {
		return nil, err
	}

	return &models.OngoingResponse{
		Day:       targetName,
		DaysAhead: days,
		Titles:    titles,
	}, nil
}
