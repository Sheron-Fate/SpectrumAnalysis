package api

import (
	"net/http"

	"colorLex/internal/app/handler"

	"github.com/gorilla/mux"
)

type Server struct{}

func NewServer() *Server { return &Server{} }

func (s *Server) Routes() {
	r := mux.NewRouter()

	// статика
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// маршруты
	r.HandleFunc("/pigments", handler.ListPigments).Methods("GET")
	r.HandleFunc("/pigment", handler.ShowPigment).Methods("GET")
	r.HandleFunc("/request/{id}", handler.ShowRequest).Methods("GET")

	// alias для главной
	r.HandleFunc("/", handler.ListPigments).Methods("GET")

	http.Handle("/", r)
}
