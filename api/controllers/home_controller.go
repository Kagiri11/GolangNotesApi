package controllers

import (
	"kagiri/GolangNotesApi/api/controllers/responses"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSONThis(w, http.StatusOK, "Hello from home")
}
