package relation

import (
	"cooking-web-app/services/ingredient"
)

func (r *RelationIngredientUser) IngredientDTO() ingredient.IngredientDTO {
	return ingredient.IngredientDTO{
		Name:          r.Ingredient.Name,
		Quantity:      r.Quantity,
		QuantityValue: r.QuantityValue,
	}
}

func (r *RelationIngredientRecipe) IngredientDTO() ingredient.IngredientDTO {
	return ingredient.IngredientDTO{
		Name:          r.Ingredient.Name,
		Quantity:      r.Quantity,
		QuantityValue: r.QuantityValue,
	}
}

func GetIngredientsDTOUser(ingredients []RelationIngredientUser) []ingredient.IngredientDTO {
	ingredientsDTO := make([]ingredient.IngredientDTO, len(ingredients))

	for i, in := range ingredients {
		ingredientsDTO[i] = in.IngredientDTO()
	}

	return ingredientsDTO
}

func GetIngredientsDTORecipe(ingredients []RelationIngredientRecipe) []ingredient.IngredientDTO {
	ingredientsDTO := make([]ingredient.IngredientDTO, len(ingredients))

	for i, in := range ingredients {
		ingredientsDTO[i] = in.IngredientDTO()
	}

	return ingredientsDTO
}
