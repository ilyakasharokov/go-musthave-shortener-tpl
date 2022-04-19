package main

import (
	"context"
	"database/sql"
	"fmt"
	"ilyakasharokov/cmd/shortener/configuration"
	"ilyakasharokov/internal/app/apiserver"
	"ilyakasharokov/internal/app/controller"
	"ilyakasharokov/internal/app/repositorydb"
	"ilyakasharokov/internal/app/worker"
	grpcshortener "ilyakasharokov/pkg/grpc"
	proto "ilyakasharokov/pkg/grpc/proto"
	"log"
	"net"
	"syscall"

	"google.golang.org/grpc"

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
	var trustedSubnet *net.IPNet = nil
	if cfg.TrustedSubnet != "" {
		_, trustedSubnet, err = net.ParseCIDR(cfg.TrustedSubnet)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	ctrl := controller.NewController(repo, cfg.BaseURL, db, wp, trustedSubnet)
	s := apiserver.New(repo, cfg.ServerAddress, trustedSubnet, db, ctrl)
	go func() {
		log.Println(s.Start(cfg.EnableHTTPS))
		cancel()
	}()

	if cfg.EnableGRPC {
		gs := grpc.NewServer()
		rungRPC(repo, &cfg, trustedSubnet, wp, db, gs, cancel, ctrl)
	}
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	select {
	case <-sigint:
		cancel()
	case <-ctx.Done():
	}
	ctxt, cancelt := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelt()
	s.Cancel(ctxt)
}

// Run gRPC server
func rungRPC(repo *repositorydb.RepositoryDB, c *configuration.Config, subnet *net.IPNet, wp *worker.WorkerPool, db *sql.DB, s *grpc.Server, stop context.CancelFunc,
	ctrl *controller.Controller) {
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		stop()
	}
	proto.RegisterShortenerServer(s, grpcshortener.New(repo, c.BaseURL, subnet, db, wp, ctrl))
	fmt.Println("gRPC server started on :3200")

	// get request from gRPC
	go func() {
		if err := s.Serve(listen); err != nil {
			stop()
		}
	}()
}
