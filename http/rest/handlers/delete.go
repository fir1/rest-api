package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (s service) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, err, 0)
			return
		}

		err = s.toDoService.Delete(r.Context(), id)
		if err != nil {
			s.respond(w, err, 0)
			return
		}
		s.respond(w, nil, http.StatusOK)
	}
}
