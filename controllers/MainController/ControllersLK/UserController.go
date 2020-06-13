package ControllersLK

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
)

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Preload("SubscriptionStatus").Preload("SubscriptionType").First(&user, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(user)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}
