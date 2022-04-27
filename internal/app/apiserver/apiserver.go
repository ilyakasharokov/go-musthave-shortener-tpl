// HTTP сервер
package apiserver

import (
	"context"
	"database/sql"
	"ilyakasharokov/internal/app/certificate"
	"ilyakasharokov/internal/app/controller"
	"ilyakasharokov/internal/app/handlers"
	"ilyakasharokov/internal/app/middlewares"
	"ilyakasharokov/internal/app/repositorydb"
	"net"
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

func New(repo *repositorydb.RepositoryDB, serverAddress string, trustedSubnet *net.IPNet, database *sql.DB, ctrl *controller.Controller) *APIServer {
	r := chi.NewRouter()
	r.Use(middlewares.GzipHandle)
	r.Use(middlewares.CookieMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	r.Post("/", handlers.CreateShort(ctrl))
	r.Post("/api/shorten", handlers.APICreateShort(ctrl))
	r.Post("/api/shorten/batch", handlers.BunchSaveJSON(ctrl))
	r.Get("/{id:[0-9a-zA-z]+}", handlers.GetShort(repo))
	r.Get("/user/urls", handlers.GetUserShorts(repo))
	r.Get("/ping", handlers.Ping(database))
	r.Delete("/api/user/urls", handlers.Delete(ctrl))
	r.Get("/api/internal/stats", handlers.Stats(repo, trustedSubnet))

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

func (s *APIServer) Start(https bool) error {
	if https {
		log.Info().Msg("Start https server on " + s.srv.Addr)
		err := certificate.Create()
		if err != nil {
			return err
		}
		return s.srv.ListenAndServeTLS("server.crt", "server.key")
	} else {
		log.Info().Msg("Start http server on " + s.srv.Addr)
		return s.srv.ListenAndServe()
	}
}

func (s *APIServer) StartTLS() error {
	log.Info().Msg("Start https server on " + s.srv.Addr)
	err := certificate.Create()
	if err != nil {
		return err
	}
	return s.srv.ListenAndServeTLS("server.crt", "server.key")
}
