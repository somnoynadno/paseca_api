package ControllerLK

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
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

type HiveCoords struct {
	CoordX *int `json:"coord_x"`
	CoordY *int `json:"coord_x"`
	HiveID uint `json:"hive_id"`
}

var SetHiveCoords = func(w http.ResponseWriter, r *http.Request) {
	coords := &HiveCoords{}

	err := json.NewDecoder(r.Body).Decode(coords)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	Hive := &models.Hive{}

	db := db.GetDB()
	err = db.First(&Hive, coords.HiveID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	Hive.CoordX = coords.CoordX
	Hive.CoordY = coords.CoordY

	err = db.Save(&Hive).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

type HiveBeeFamily struct {
	BeeFamilyID *uint `json:"bee_family_id"`
	HiveID      *uint `json:"hive_id"`
}

var SetHiveBeeFamily = func(w http.ResponseWriter, r *http.Request) {
	hiveBeeFamily := &HiveBeeFamily{}

	err := json.NewDecoder(r.Body).Decode(hiveBeeFamily)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	Hive := &models.Hive{}
	BeeFamily := &models.BeeFamily{}

	err = db.First(&Hive, hiveBeeFamily.HiveID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	Hive.BeeFamilyID = hiveBeeFamily.BeeFamilyID
	err = db.Save(&Hive).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.First(&BeeFamily, hiveBeeFamily.BeeFamilyID).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	BeeFamily.HiveID = hiveBeeFamily.HiveID
	err = db.Save(&BeeFamily).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	u.Respond(w, u.Message(true, "OK"))
}

