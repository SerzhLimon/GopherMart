package main

import (
	"os"
	"os/signal"
	"syscall"

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
		logrus.Warn(err)
	}

	logrus.Info("Running migrations...")
	err = migrations.Up(psql)
	if err != nil {
		logrus.Warn(err)
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s.Run()
		quit <- syscall.SIGTERM
	}()

	sig := <-quit
	logrus.Infof("Received signal: %v", sig)
}
