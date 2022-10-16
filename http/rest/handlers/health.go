package handlers

import (
	"net/http"
)

func (s service) Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, "OK", http.StatusOK)
	}
}
