package main

import (
	"context"
	"ilyakasharokov/internal/app/apiserver"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	s := apiserver.New()
	go func() {
		log.Fatal(s.Start(ctx))
		cancel()
	}()
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	select {
	case <-sigint:
		cancel()
	case <-ctx.Done():
	}
	s.Shutdown()
}
