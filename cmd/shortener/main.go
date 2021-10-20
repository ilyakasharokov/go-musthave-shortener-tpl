package main

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"ilyakasharokov/cmd/shortener/configuration"
	"ilyakasharokov/internal/app/apiserver"
	"ilyakasharokov/internal/app/dbservice"
	"ilyakasharokov/internal/app/repositorydb"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg := configuration.New()
	db, err := sql.Open("postgres", cfg.Database)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	dbservice.SetupDatabase(db, ctx)
	repo := repositorydb.New(db)
	s := apiserver.New(repo, cfg.ServerAddress, cfg.BaseURL, db)
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
