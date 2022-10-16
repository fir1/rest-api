package handlers

import (
	"github.com/fir1/rest-api/internal/todo/model"
	toDoService "github.com/fir1/rest-api/internal/todo/service"
	"net/http"
)

func (s service) Create() http.HandlerFunc {
	type request struct {
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
	}

	type response struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := s.decode(r, &req)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		id, err := s.toDoService.Create(r.Context(), toDoService.CreateParams{
			Name:        req.Name,
			Description: req.Description,
			Status:      req.Status,
		})
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{ID: id}, http.StatusOK)
	}
}
