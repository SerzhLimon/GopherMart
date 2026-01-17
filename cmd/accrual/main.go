package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/SerzhLimon/GopherMart/internal/config_accrual"
	db "github.com/SerzhLimon/GopherMart/internal/config_accrual/db"
	server "github.com/SerzhLimon/GopherMart/internal/server_accrual"
	migrations "github.com/SerzhLimon/GopherMart/migrations_accrual"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalln(err)
	}
	psql, err := db.InitPostgresClient(cfg)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer psql.Close()

	logrus.Info("Running migrations...")
	err = migrations.Up(psql)
	if err != nil {
		logrus.Fatalln(err)
	} else {
		logrus.Info("Migrations applied successfully")
	}
	defer func() {
		migrations.Down(psql)
		logrus.Info("Migrations down")
	}()

	s, err := server.NewServer(cfg, psql)
	if err != nil {
		logrus.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go s.RunEngine(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	serverErr := make(chan error, 1)
	go func() {
		if err := s.Run(); err != nil {
			serverErr <- err
		}
	}()

	select {
	case <-quit:
		logrus.Info("Shutdown signal received")
	case err := <-serverErr:
		logrus.WithError(err).Error("Server error occurred")
	}

	cancel()                           
	time.Sleep(500 * time.Millisecond)

	logrus.Info("Shutting down...")
}
