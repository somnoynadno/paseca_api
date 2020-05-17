package ControllerLK

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
)

var GetUser = func(w http.ResponseWriter, r *http.Request) {
	var entities models.User
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Preload("SubscriptionType").Preload("SubscriptionStatus").
		Where("id = ?", id).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}
