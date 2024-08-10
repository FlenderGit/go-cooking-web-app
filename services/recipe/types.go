package recipe

import (
	"cooking-web-app/services/ingredient"
	"cooking-web-app/services/relation"
)

type Recipe struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Ingredients []relation.RelationIngredientRecipe
	NbPeople    uint
	Steps       string
}

type RecipeDTO struct {
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Ingredients []ingredient.IngredientDTO `json:"ingredients"`
	NbPeople    uint                       `json:"number_people"`
	Steps       string                     `json:"steps"`
}

type RecipePost struct {
	Name        string                            `json:"name"`
	Description string                            `json:"description"`
	Ingredients []ingredient.IngredientRecipePost `json:"ingredients"`
	NbPeople    uint                              `json:"number_people"`
	Steps       string                            `json:"steps"`
}

func (r *Recipe) ToRecipeSimple() *RecipeDTO {
	return &RecipeDTO{
		Name:        r.Name,
		Description: r.Description,
		Ingredients: relation.GetIngredientsDTORecipe(r.Ingredients),
		NbPeople:    r.NbPeople,
		Steps:       r.Steps,
	}
}
