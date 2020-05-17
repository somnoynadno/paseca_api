package ControllerLK

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
)

// TODO: add order by on different columns
// TODO: add pagination
var GetUsersHoneySales = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.HoneySale
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Preload("HoneyType").Preload("BeeFarm").
		Where("user_id = ?", id).Order("updated_at desc").Find(&entities).Error

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
