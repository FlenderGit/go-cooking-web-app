package user

import (
	"cooking-web-app/services/auth"
	"net/http"

	"gorm.io/gorm"
)

type IStoreMe interface {
	GetMe() (*User, error)
}

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetMe(r *http.Request) (*User, error) {

	user_jwt, err := auth.GetUserFromRequest(r)
	if err != nil {
		return nil, err
	}

	var user User
	err = s.db.Preload("Ingredients.Ingredient").First(&user, user_jwt.ID).Error
	return &user, err
}
