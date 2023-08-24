package server

import (
	"log"
	"net/http"

	"github.com/amaghzaz-y/torrex/internal/streamer"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	router *chi.Mux
}

func DefaultServer() *Server {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	return &Server{
		router,
	}
}

func (s *Server) Start() error {
	handler := streamer.NewStreamer().Test()
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	s.router.HandleFunc("/*", handler)
	return http.ListenAndServe("127.0.0.1:4000", s.router)
}
