package usecase

import (
	"database/sql"
	"fmt"

	repo "github.com/SerzhLimon/GopherMart/internal/repository"
)

type UseCase interface {
}

type Usecase struct {
	repo repo.Repository
}

func NewService(db *sql.DB) (UseCase, error) {
	repo, err := repo.NewRepo(db)
	if err != nil {
		return nil, fmt.Errorf("fail to init repo")
	}
	
	return &Usecase{
		repo: repo,
	}, nil
}
