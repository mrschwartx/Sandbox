package store

import (
	"context"
	"errors"
	"sync"
)

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductStore struct {
	mu       sync.RWMutex
	Products map[int64]*Product
}

func NewProductStore() *ProductStore {
	store := &ProductStore{
		Products: make(map[int64]*Product),
	}

	store.Products[1] = &Product{ID: 1, Name: "Laptop", Price: 1500, Stock: 5}
	store.Products[2] = &Product{ID: 2, Name: "Keyboard", Price: 80, Stock: 20}
	store.Products[3] = &Product{ID: 3, Name: "Mouse", Price: 25, Stock: 100}

	return store
}

func (s *ProductStore) GetByName(ctx context.Context, name string) (*Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, p := range s.Products {
		if p.Name == name {
			return p, nil
		}
	}
	return nil, errors.New("product not found")
}

func (s *ProductStore) GetAll(ctx context.Context) ([]Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]Product, 0, len(s.Products))
	for _, p := range s.Products {
		result = append(result, *p)
	}
	return result, nil
}
