package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kagiri/GolangNotesApi/models"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize() {
	var err error
	dsn := "root:Positive11@tcp(127.0.0.1:3306)/pleya?charset=utf8mb4&parseTime=True&loc=Local"
	s.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Encountered challenge when connecting to DB: %v", err)
		return
	}

	err = s.DB.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		log.Fatalf("Encountered challenge when auto migrating: %v", err)
		return
	}
	s.Router = mux.NewRouter()
	s.InitializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
