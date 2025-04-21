package auth

import (
	"github.com/tranvinh21/fastext-be-go/db/schema"
	"gorm.io/gorm"
)

type AuthStore struct {
	db *gorm.DB
}

func NewAuthStore(db *gorm.DB) *AuthStore {
	return &AuthStore{db: db}
}

func (s *AuthStore) CreateUser(user *schema.User) error {
	return s.db.Create(user).Error
}

func (s *AuthStore) GetUserByEmail(email string) (*schema.User, error) {
	var user schema.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *AuthStore) GetUserByUsername(username string) (*schema.User, error) {
	var user schema.User
	if err := s.db.Where("name = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
