// HTTP сервер
package apiserver

import (
	"context"
	"database/sql"
	"ilyakasharokov/internal/app/handlers"
	"ilyakasharokov/internal/app/middlewares"
	"ilyakasharokov/internal/app/repositorydb"
	"ilyakasharokov/internal/app/worker"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	repo repositorydb.RepositoryDB
	srv  *http.Server
	db   *sql.DB
}

func New(repo *repositorydb.RepositoryDB, serverAddress string, baseURL string, database *sql.DB, wp *worker.WorkerPool) *APIServer {
	r := chi.NewRouter()
	r.Use(middlewares.GzipHandle)
	r.Use(middlewares.CookieMiddleware)
	r.Post("/", handlers.CreateShort(repo, baseURL))
	r.Post("/api/shorten", handlers.APICreateShort(repo, baseURL))
	r.Post("/api/shorten/batch", handlers.BunchSaveJSON(repo, baseURL))
	r.Get("/{id:[0-9a-zA-z]+}", handlers.GetShort(repo))
	r.Get("/user/urls", handlers.GetUserShorts(repo))
	r.Get("/ping", handlers.Ping(database))
	r.Delete("/api/user/urls", handlers.Delete(repo, wp))

	r.Mount("/debug/", middleware.Profiler())

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
	log.Info().Msg("Start http server on " + s.srv.Addr)
	err := s.srv.ListenAndServe()
	return err
}
