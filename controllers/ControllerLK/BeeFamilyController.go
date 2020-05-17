package ControllerLK

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	u "paseca/utils"
)

type BeeFamilyShort struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	BeeFarmName string  `json:"bee_farm_name"`
}

var GetUsersBeeFamilies = func(w http.ResponseWriter, r *http.Request) {
	var entities []BeeFamilyShort
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Table("bee_families").
		Joins("join bee_farms on bee_farms.id = bee_families.bee_farm_id").
		Joins("join users on users.id = bee_farms.user_id").
		Select("bee_families.id, bee_families.name, bee_farms.name as bee_farm_name").
		Where("users.id = ?", id).Scan(&entities).Error

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
