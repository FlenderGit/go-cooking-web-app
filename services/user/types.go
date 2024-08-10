package user

import (
	"cooking-web-app/services/ingredient"
	"cooking-web-app/services/relation"
)

type User struct {
	ID           int `gorm:"primaryKey"`
	Login        string
	HashPassword string
	Email        string
	Ingredients  []relation.RelationIngredientUser
}

type UserSimpleDTO struct {
	Login       string                     `json:"user"`
	Email       string                     `json:"email"`
	Ingredients []ingredient.IngredientDTO `json:"ingredients"`
}

func (u *User) UserSimpleDTO() *UserSimpleDTO {
	return &UserSimpleDTO{
		Login:       u.Login,
		Email:       u.Email,
		Ingredients: relation.GetIngredientsDTOUser(u.Ingredients),
	}
}
