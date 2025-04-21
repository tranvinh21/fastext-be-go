package user

import (
	"github.com/tranvinh21/fastext-be-go/db/schema"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) GetUsers() ([]schema.User, error) {
	var users []schema.User
	err := s.db.Find(&users).Error
	return users, err
}
