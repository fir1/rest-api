package handlers

import (
	toDoRepo "github.com/fir1/rest-api/internal/todo/repository"
	toDoService "github.com/fir1/rest-api/internal/todo/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger      *logrus.Logger
	router      *mux.Router
	toDoService toDoService.Service
}

func newHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:      lg,
		toDoService: toDoService.NewService(toDoRepo.NewRepository(db)),
	}
}
