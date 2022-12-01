package controllers

import (
	"encoding/json"
	"io/ioutil"
	"kagiri/GolangNotesApi/api/controllers/responses"
	"kagiri/GolangNotesApi/models"
	"net/http"
)

func (s *Server) SignUpUser(w http.ResponseWriter, r *http.Request) {
	// get raw body from api
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadGateway, err)
		return
	}
	// unmarshal that json body to a struct
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// save that user to db
	createdUser, err := user.CreateUser(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusForbidden, err)
		return
	}

	// reply to the api
	responses.JSONThis(w, http.StatusCreated, createdUser)
}

func (s *Server) SignInUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
