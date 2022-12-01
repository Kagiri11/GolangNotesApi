package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"kagiri/GolangNotesApi/api/controllers/responses"
	"kagiri/GolangNotesApi/models"
	"net/http"
	"strconv"
)

func (s *Server) CreateNote(w http.ResponseWriter, r *http.Request) {
	// get raw json body from server and check for errors
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadGateway, err)
		return
	}
	// unmarshall the json to go note struct we will need a go model for this check for unmarshalling error
	note := models.Note{}
	err = json.Unmarshal(body, &note)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// save the note struct to the db..use db provided by the server
	createdNote, err := note.CreateNote(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusConflict, err)
		return
	}

	responses.JSONThis(w, http.StatusCreated, &createdNote)
}

func (s *Server) GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	note := models.Note{}
	noteFromDB, err := note.GetNote(s.DB, id)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSONThis(w, http.StatusFound, noteFromDB)
}

func (s *Server) GetNotes(w http.Request, r *http.Request) {

}

func (s *Server) UpdateNote(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) DeleteNote(w http.ResponseWriter, r *http.Request) {

}
