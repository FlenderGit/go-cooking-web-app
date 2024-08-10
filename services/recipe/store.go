package recipe

import (
	"cooking-web-app/services/relation"

	"gorm.io/gorm"
)

type IStoreMe interface {
	GetRecipeById(int) (*Recipe, error)
	GetRandomRecipe() (*Recipe, error)
	CreateRecipe(*RecipeDTO) error
}

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetRecipeById(id int) (*Recipe, error) {
	var recipe Recipe
	err := s.db.Preload("Ingredients.Ingredient").First(&recipe, id).Error

	return &recipe, err
}

func (s *Store) GetRandomRecipe() (*Recipe, error) {
	var recipe Recipe
	err := s.db.Preload("Ingredients.Ingredient").Order("RANDOM()").First(&recipe).Error

	return &recipe, err
}

func (s *Store) CreateRecipe(recipe *RecipePost) error {
	recipe_db := Recipe{
		Name:        recipe.Name,
		Description: recipe.Description,
		NbPeople:    recipe.NbPeople,
		Steps:       recipe.Steps,
	}

	err := s.db.Create(&recipe_db).Error
	if err != nil {
		return err
	}

	for _, ingredient := range recipe.Ingredients {
		ingredient_db := relation.RelationIngredientRecipe{
			RecipeID:      recipe_db.ID,
			IngredientID:  ingredient.ID,
			Quantity:      ingredient.Quantity,
			QuantityValue: ingredient.QuantityValue,
		}
		err = s.db.Create(&ingredient_db).Error
		if err != nil {
			return err
		}
	}

	return nil
}
