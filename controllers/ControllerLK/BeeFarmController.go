package ControllerLK

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
)

type BeeFarm struct {
	models.BaseModel
	Name          string       `json:"name"`
	Location      *string      `json:"location"`
	BeeFarmTypeID uint         `json:"bee_farm_type_id"`
	BeeFarmType   BeeFarmType  `json:"bee_farm_type"`
	BeeFarmSizeID uint         `json:"bee_farm_size_id"`
	BeeFarmSize   BeeFarmSize  `json:"bee_farm_size"`
}

var GetUserBeeFarms = func(w http.ResponseWriter, r *http.Request) {
	var entities []BeeFarm
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Where("user_id = ?", id).Find(&entities).Error

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

var GetBeeFarmByID = func(w http.ResponseWriter, r *http.Request) {
	var BeeFarm models.BeeFarm

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("Reminders").Preload("BeeFamilies").
		Preload("Hives").Where("id = ?", id).Find(&BeeFarm).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(BeeFarm)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}



