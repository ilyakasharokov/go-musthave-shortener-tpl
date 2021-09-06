package apiserver

import (
	"github.com/go-chi/chi/v5"
	"ilyakasharokov/internal/app/handlers"
	"ilyakasharokov/internal/app/repository"
	"log"
	"net/http"
)

type APIServer struct {
	BindAddr string
	repo     *repository.Repository
}

func New() *APIServer {
	return &APIServer{
		BindAddr: ":8080",
	}
}

func (s *APIServer) Start() error {
	s.repo = repository.New()
	r := chi.NewRouter()
	r.Post("/", handlers.CreateShort(s.repo))
	r.Get("/{id:[0-9a-z]+}", handlers.GetShort(s.repo))
	err := http.ListenAndServe(s.BindAddr, r)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
