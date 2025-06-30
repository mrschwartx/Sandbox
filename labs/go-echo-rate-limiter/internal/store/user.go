package store

import (
	"context"
	"errors"
	"sync"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"` // Note: hashed in real applications
}

type UserStore struct {
	mu    sync.RWMutex
	Users map[int64]*User
}

func NewStoreUser() *UserStore {
	store := &UserStore{
		Users: make(map[int64]*User),
	}

	store.Users[1] = &User{ID: 1, Email: "alice@example.com", Password: "hashedpass1"}
	store.Users[2] = &User{ID: 2, Email: "bob@example.com", Password: "hashedpass2"}

	return store
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, user := range s.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}
