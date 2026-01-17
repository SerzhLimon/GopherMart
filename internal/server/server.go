package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/SerzhLimon/GopherMart/internal/config"
	uc "github.com/SerzhLimon/GopherMart/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	cfg  *config.Config
	core *chi.Mux
	uc   uc.UseCase
}

func NewServer(cfg  *config.Config, db *sql.DB) (*Server, error) {
	uc, err := uc.NewService(db)
	if err != nil {
		return nil, err
	}
	server := &Server{
		core: chi.NewRouter(),
		uc:   uc,
		cfg: cfg,
	}
	server.route()
	return server, nil
}

func (s *Server) route() {
	// s.core.Use(handLogger)

	// s.core.Post("/", s.SetURL)
	// s.core.Post("/api/shorten", s.SetURLJson)
}

func (s *Server) Run() {
	logrus.Infof("server started with params: host - %s", s.cfg.Opts.ServerHost)
	if err := http.ListenAndServe(s.cfg.Opts.ServerHost, s.core); err != nil {
		log.Fatalln(err)
	}
}
