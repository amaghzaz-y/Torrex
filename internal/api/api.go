package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type API struct {
	router *chi.Mux
}

func New() *API {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	return &API{
		r,
	}
}
func (a *API) Handle(path string, handler http.HandlerFunc) {
	a.router.HandleFunc(path, handler)
}

func (a *API) Start() error {
	a.router.HandleFunc("/search/{query}", SearchHandler)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	return http.ListenAndServe("127.0.0.1:4000", a.router)
}
