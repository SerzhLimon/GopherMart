package main

import (
	"github.com/SerzhLimon/GopherMart/internal/config"
	"github.com/SerzhLimon/GopherMart/internal/config/db"
	"github.com/SerzhLimon/GopherMart/internal/server"
	"github.com/SerzhLimon/GopherMart/migrations"
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
		//save data
		// migrations.Down(psql)
		// logrus.Info("Migrations down")
	}()

	s, err := server.NewServer(cfg, psql)
	if err != nil {
		logrus.Fatalln(err)
	}
	s.Run()
}
