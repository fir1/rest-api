package handlers

import (
	"errors"
	"github.com/fir1/rest-api/internal/todo/model"
	"github.com/fir1/rest-api/pkg/erru"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func (s service) Get() http.HandlerFunc {
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

		getResponse, err := s.toDoService.Get(r.Context(), id)
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, response{
			ID:          getResponse.ID,
			Name:        getResponse.Name,
			Description: getResponse.Description,
			Status:      getResponse.Status,
			CreatedOn:   getResponse.CreatedOn,
			UpdatedOn:   getResponse.UpdatedOn,
		}, http.StatusOK)
	}
}
