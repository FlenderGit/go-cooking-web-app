package user

import (
	"cooking-web-app/services/auth"
	"cooking-web-app/utils"
	"net/http"

	"gorm.io/gorm"
)

func GetMeRouter(db *gorm.DB) *utils.Router {

	me_store := NewStore(db)
	me_router := utils.NewRouter()

	me_router.Get("/{$}", func(w http.ResponseWriter, r *http.Request) {
		user, err := me_store.GetMe(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		utils.Encode(w, r, http.StatusOK, user.UserSimpleDTO())
	}, auth.IsAuthenticated)

	return me_router
}
