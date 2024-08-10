package relation

import "cooking-web-app/services/ingredient"

type RelationIngredientRecipe struct {
	RecipeID      uint `gorm:"primaryKey"`
	IngredientID  uint `gorm:"primaryKey"`
	Ingredient    ingredient.Ingredient
	Quantity      float64
	QuantityValue string
}

type RelationIngredientUser struct {
	UserID        uint `gorm:"primaryKey"`
	IngredientID  uint `gorm:"primaryKey"`
	Ingredient    ingredient.Ingredient
	Quantity      float64
	QuantityValue string
}
