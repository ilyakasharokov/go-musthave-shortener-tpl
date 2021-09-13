package apiserver

import (
	"context"
	"github.com/go-chi/chi/v5"
	"ilyakasharokov/internal/app/handlers"
	"ilyakasharokov/internal/app/repository"
	"net/http"
)

type APIServer struct {
	repo repository.RepoModel
	srv  *http.Server
}

func New(repo repository.RepoModel, addr string) *APIServer {
	r := chi.NewRouter()
	r.Post("/", handlers.CreateShort(repo))
	r.Get("/{id:[0-9a-z]+}", handlers.GetShort(repo))
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	return &APIServer{
		repo: repo,
		srv:  srv,
	}
}

func (s *APIServer) Cancel(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *APIServer) Start() error {
	err := s.srv.ListenAndServe()
	return err
}
