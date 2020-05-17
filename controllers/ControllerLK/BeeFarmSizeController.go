package ControllerLK

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	u "paseca/utils"
)

type BeeFarmSize struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	MaxX        int     `json:"max_x"`
	MaxY        int     `json:"max_y"`
}

var GetBeeFarmSizes = func(w http.ResponseWriter, r *http.Request) {
	var entities []BeeFarmSize
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
