package repository_accrual

import (
	"context"
	"database/sql"
	"fmt"

	m "github.com/SerzhLimon/GopherMart/internal/models_accrual"
)

type Repository interface {
	CreateOrder(req *m.CreateOrderRequest) error
	IsExistOrder(orderID string) (bool, error)
}

type PgStorage struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) (Repository, error) {
	if db == nil {
		return nil, fmt.Errorf("db not init")
	}
	return &PgStorage{
		db: db,
	}, nil
}

func (r *PgStorage) IsExistOrder(orderID string) (bool, error) {
	var exist bool
	err := r.db.QueryRow(queryExistOrderID, orderID).Scan(&exist)
	if err != nil {
		return exist, fmt.Errorf("repo.IsExistOrder() %w", err)
	}

	return exist, nil
}

func (r *PgStorage) CreateOrder(req *m.CreateOrderRequest) error {
	tx, err := r.db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, good := range req.Goods {
		_, err := r.db.Exec(queryCreateOrder, *req.Order, good.Description, good.Price)
		if err != nil {
			return fmt.Errorf("repo.CreateOrder() %w", err)
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
