package apiserver

import (
	"context"
	"github.com/go-chi/chi/v5"
	"ilyakasharokov/internal/app/handlers"
	"ilyakasharokov/internal/app/repository"
	"net/http"
)

type APIServer struct {
	repo repository.Repository
	srv  *http.Server
}

func New(repo *repository.Repository, serverAddress string, baseURL string) *APIServer {
	r := chi.NewRouter()
	r.Post("/", handlers.CreateShort(repo, baseURL))
	r.Post("/api/shorten", handlers.APICreateShort(repo, baseURL))
	r.Get("/{id:[0-9a-z]+}", handlers.GetShort(repo))
	srv := &http.Server{
		Addr:    serverAddress,
		Handler: r,
	}
	return &APIServer{
		repo: *repo,
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
