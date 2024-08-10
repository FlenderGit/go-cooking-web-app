package api

import (
	"cooking-web-app/services/recipe"
	"cooking-web-app/services/user"
	"cooking-web-app/utils"

	"gorm.io/gorm"
)

func GetApi(db *gorm.DB) *utils.Router {
	api_router := utils.NewRouter()

	user.GetMeRouter(db).MountOnRouter(api_router, "/me")
	recipe.GetRecipeRouter(db).MountOnRouter(api_router, "/recipe")

	return api_router
}
