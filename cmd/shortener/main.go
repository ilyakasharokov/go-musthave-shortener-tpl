package main

import (
	"context"
	"database/sql"
	"ilyakasharokov/cmd/shortener/configuration"
	"ilyakasharokov/internal/app/apiserver"
	"ilyakasharokov/internal/app/repositorydb"
	"ilyakasharokov/internal/app/worker"
	"log"

	_ "github.com/lib/pq"

	"os"
	"os/signal"
	"time"
)

var (
	buildVersion = "N/A"
	buildDate    = "N/A"
	buildCommit  = "N/A"
)

func main() {
	log.Printf("Build version: %v\n", buildVersion)
	log.Printf("Build date: %v\n", buildDate)
	log.Printf("Build commit: %v\n", buildCommit)

	ctx, cancel := context.WithCancel(context.Background())
	cfg := configuration.New()
	db, err := sql.Open("postgres", cfg.Database)
	if err != nil {
		return
	}
	defer db.Close()
	repo := repositorydb.New(db)
	wp := worker.New(5, 5)
	go wp.Run(ctx)
	s := apiserver.New(repo, cfg.ServerAddress, cfg.BaseURL, db, wp)
	go func() {
		log.Println(s.Start())
		cancel()
	}()
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	select {
	case <-sigint:
		cancel()
	case <-ctx.Done():
	}
	ctxt, cancelt := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelt()
	s.Cancel(ctxt)
}
