package service

import (
	"github.com/fir1/rest-api/internal/todo/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repo: r,
	}
}
