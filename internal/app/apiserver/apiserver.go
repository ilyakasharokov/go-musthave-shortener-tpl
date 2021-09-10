package apiserver

import (
	"context"
	"github.com/go-chi/chi/v5"
	"ilyakasharokov/internal/app/handlers"
	"ilyakasharokov/internal/app/repository"
	"net/http"
)

type APIServer struct {
	BindAddr string
	repo     repository.RepoModel
	context  context.Context
	srv      *http.Server
}

func New() *APIServer {
	return &APIServer{
		BindAddr: ":8080",
	}
}

func (s *APIServer) Shutdown() error {
	return s.srv.Shutdown(s.context)
}

func (s *APIServer) Start(ctx context.Context) error {
	s.repo = repository.New()
	s.context = ctx
	r := chi.NewRouter()
	s.srv = &http.Server{
		Addr:    s.BindAddr,
		Handler: r,
	}
	r.Post("/", handlers.CreateShort(s.repo))
	r.Get("/{id:[0-9a-z]+}", handlers.GetShort(s.repo))
	err := s.srv.ListenAndServe()
	return err
}
