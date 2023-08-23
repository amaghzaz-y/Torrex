package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

func DefaultServer() *Server {
	router := chi.NewRouter()
	return &Server{
		router,
	}
}

func (s *Server) Start() {
	http.ListenAndServe(":3000", s.router)
}
