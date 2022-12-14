package controllers

import "kagiri/GolangNotesApi/api/middlewares"

func (s *Server) InitializeRoutes() {
	s.Router.Handle("/home", middlewares.SetJSONMiddleware(s.Home)).Methods("GET")

	// User Routes
	s.Router.Handle("/user", middlewares.SetJSONMiddleware(s.SignUpUser)).Methods("POST")
	s.Router.Handle("/user/{id}", middlewares.SetJSONMiddleware(s.GetUser)).Methods("GET")
	s.Router.Handle("/users", middlewares.SetJSONMiddleware(s.GetUsers)).Methods("GET")
	s.Router.Handle("/user/{id}", middlewares.SetJSONMiddleware(s.DeleteUser)).Methods("DELETE")

	// Note Routes
	s.Router.Handle("/note", middlewares.SetJSONMiddleware(s.CreateNote)).Methods("POST")
	s.Router.Handle("/note/{id}", middlewares.SetJSONMiddleware(s.GetNote)).Methods("GET")
	s.Router.Handle("/notes", middlewares.SetJSONMiddleware(s.GetNotes)).Methods("GET")
	s.Router.Handle("/note/{id}", middlewares.SetJSONMiddleware(s.DeleteNote)).Methods("DELETE")
}
