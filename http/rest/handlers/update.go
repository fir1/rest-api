package handlers

import (
	"errors"
	"github.com/fir1/rest-api/internal/todo/model"
	toDoService "github.com/fir1/rest-api/internal/todo/service"
	"github.com/fir1/rest-api/pkg/erru"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (s service) Update() http.HandlerFunc {
	type request struct {
		Name        *string       `json:"name"`
		Description *string       `json:"description"`
		Status      *model.Status `json:"status"`
	}

	type response struct {
		ID          int          `json:"id"`
		Name        string       `json:"name"`
		Description string       `json:"description"`
		Status      model.Status `json:"status"`
		CreatedOn   time.Time    `json:"created_on"`
		UpdatedOn   *time.Time   `json:"updated_on,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, erru.ErrArgument{
				Wrapped: errors.New("valid id must provide in path"),
			}, 0)
			return
		}

		req := request{}
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err = s.decode(r, &req)
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		err = s.toDoService.Update(r.Context(), toDoService.UpdateParams{
			ID:          id,
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
