package main

import (
	"context"
	"ilyakasharokov/cmd/shortener/configuration"
	"ilyakasharokov/internal/app/apiserver"
	"ilyakasharokov/internal/app/repository"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg := configuration.New()
	repo := repository.New(cfg.FileStoragePath)
	s := apiserver.New(repo, cfg.ServerAddress, cfg.BaseURL)
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
	ctxt, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.Cancel(ctxt)
}
