package recipe

import (
	"cooking-web-app/utils"
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	Store *Store
}

func GetRecipeRouter(db *gorm.DB) *utils.Router {

	handler := &Handler{
		Store: NewStore(db),
	}
	recipe_router := utils.NewRouter()

	recipe_router.Get("/{id}", handler.handleGetRecipe)
	recipe_router.Get("/random", handler.handleGetRandomRecipe)
	recipe_router.Post("/{$}", handler.handleCreateRecipe)

	return recipe_router
}

func (h *Handler) handleGetRecipe(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetIdFromPath(r)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	recipe, err := h.Store.GetRecipeById(id)
	if err != nil {
		http.Error(w, "Recipe not found", http.StatusNotFound)
		return
	}

	utils.Encode(w, r, http.StatusOK, recipe.ToRecipeSimple())
}

func (h *Handler) handleGetRandomRecipe(w http.ResponseWriter, r *http.Request) {
	recipe, err := h.Store.GetRandomRecipe()
	if err != nil {
		http.Error(w, "No recipe found", http.StatusNotFound)
		return
	}

	utils.Encode(w, r, http.StatusOK, recipe.ToRecipeSimple())
}

func (h *Handler) handleCreateRecipe(w http.ResponseWriter, r *http.Request) {
	recipe, err := utils.Decode[RecipePost](r)
	if err != nil {
		http.Error(w, "Invalid recipe", http.StatusBadRequest)
		return
	}

	err = h.Store.CreateRecipe(&recipe)
	if err != nil {
		http.Error(w, "Failed to create recipe", http.StatusInternalServerError)
		return
	}

	utils.Encode(w, r, http.StatusCreated, "Recipe created")
}
