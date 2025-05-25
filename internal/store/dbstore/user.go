package dbstore

import (
	"database/sql"
	"goth/internal/hash"
	"goth/internal/store"

)

type UserStore struct {
	db           *sql.DB
	passwordhash hash.PasswordHash
}

type NewUserStoreParams struct {
	DB           *sql.DB
	PasswordHash hash.PasswordHash
}

func NewUserStore(params NewUserStoreParams) *UserStore {
	return &UserStore{
		db:           params.DB,
		passwordhash: params.PasswordHash,
	}
}

func (s *UserStore) CreateUser(email string, password string) error {

	// hashedPassword, err := s.passwordhash.GenerateFromPassword(password)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (s *UserStore) GetUser(email string) (*store.User, error) {
	return &store.User{}, nil 
}
