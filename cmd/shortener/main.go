package main

import (
	"context"
	"ilyakasharokov/cmd/shortener/configuration"
	"ilyakasharokov/internal/app/apiserver"
	"ilyakasharokov/internal/app/repository"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg := configuration.New()
	repo := repository.New(cfg)
	s := apiserver.New(repo, cfg)
	go func() {
		log.Fatal(s.Start())
		cancel()
	}()
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	select {
	case <-sigint:
		cancel()
	case <-ctx.Done():
	}
	s.Cancel(ctx)
}
