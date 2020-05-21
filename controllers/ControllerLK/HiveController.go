package ControllerLK

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
)

type HiveCreateModel struct {
	Name             string   `json:"name"`
	HiveFrameTypeID  uint     `json:"hive_frame_type_id"`
	HiveFormatID     uint     `json:"hive_format_id"`
	BeeFarmID        uint     `json:"bee_farm_id"`
}

var CreateHive = func(w http.ResponseWriter, r *http.Request) {
	CreateModel := &HiveCreateModel{}

	err := json.NewDecoder(r.Body).Decode(CreateModel)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	Model := models.Hive{BeeFarmID: CreateModel.BeeFarmID,
		Name: CreateModel.Name, HiveFormatID: CreateModel.HiveFormatID,
		HiveFrameTypeID: CreateModel.HiveFrameTypeID,
	}

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}
