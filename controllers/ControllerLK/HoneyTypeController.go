package ControllerLK

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	u "paseca/utils"
)

type HoneyType struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	BasePrice   int     `json:"base_price"`
}

var GetHoneyTypes = func(w http.ResponseWriter, r *http.Request) {
	var entities []HoneyType
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Where("is_custom = false or (is_custom = true and creator_id = ?)", id).Find(&entities).Error

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
