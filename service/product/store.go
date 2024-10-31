package product

import (
	"database/sql"

	"github.com/yann-fk-21/todo-platform/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	return nil, nil
}