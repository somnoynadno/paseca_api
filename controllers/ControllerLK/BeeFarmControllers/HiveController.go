package BeeFarmControllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
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
	id := u.GetUserIDFromRequest(r)

	err := json.NewDecoder(r.Body).Decode(CreateModel)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	Model := models.Hive{BeeFarmID: CreateModel.BeeFarmID,
		Name: CreateModel.Name, HiveFormatID: CreateModel.HiveFormatID,
		HiveFrameTypeID: CreateModel.HiveFrameTypeID,
	}
	Model.UserID = id

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

type HiveCoords struct {
	CoordX    *int `json:"coord_x"`
	CoordY    *int `json:"coord_y"`
	BeeFarmID uint `json:"bee_farm_id"`
	HiveID    uint `json:"hive_id"`
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
	Hive.BeeFarmID = coords.BeeFarmID

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

	if hiveBeeFamily.HiveID != nil {
		err = db.Model(models.Hive{}).First(&Hive, *hiveBeeFamily.HiveID).Error

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

		if hiveBeeFamily.BeeFamilyID == nil {
			err = db.Where("hive_id = ?", hiveBeeFamily.HiveID).First(&BeeFamily).Error

			if err != nil {
				if err == gorm.ErrRecordNotFound {
					u.HandleNotFound(w)
				} else {
					u.HandleBadRequest(w, err)
				}
				return
			}

			BeeFamily.HiveID = nil
			err = db.Save(&BeeFamily).Error

			if err != nil {
				u.HandleBadRequest(w, err)
				return
			}
		}
	}

	if hiveBeeFamily.BeeFamilyID != nil {
		err = db.First(&BeeFamily, *hiveBeeFamily.BeeFamilyID).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				u.HandleNotFound(w)
			} else {
				u.HandleBadRequest(w, err)
			}
			return
		}

		BeeFamily.HiveID = hiveBeeFamily.HiveID
		BeeFamily.BeeFarmID = Hive.BeeFarmID
		err = db.Save(&BeeFamily).Error

		if err != nil {
			u.HandleBadRequest(w, err)
			return
		}

		if hiveBeeFamily.HiveID == nil {
			err = db.Where("bee_family_id = ?", hiveBeeFamily.BeeFamilyID).First(&Hive).Error

			if err != nil {
				if err == gorm.ErrRecordNotFound {
					u.HandleNotFound(w)
				} else {
					u.HandleBadRequest(w, err)
				}
				return
			}

			Hive.BeeFamilyID = nil
			err = db.Save(&Hive).Error

			if err != nil {
				u.HandleBadRequest(w, err)
				return
			}
		}
	}

	u.Respond(w, u.Message(true, "OK"))
}

type HiveShort struct {
	ID                uint    `json:"id"`
	Name              string  `json:"name"`
	BeeFarmName       string  `json:"bee_farm_name"`
	HiveFrameTypeName string  `json:"hive_frame_type_name"`
	HiveFormatName    string  `json:"hive_format_name"`
	HiveFormatSize    *int    `json:"hive_format_size"`
}

var GetUsersFreeHives = func(w http.ResponseWriter, r *http.Request) {
	var entities []HiveShort
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Table("hives").
		Joins("join bee_farms on bee_farms.id = hives.bee_farm_id").
		Joins("join hive_formats on hive_formats.id = hives.hive_format_id").
		Joins("join hive_frame_types on hive_frame_types.id = hives.hive_frame_type_id").
		Select("hives.id, hives.name, bee_farms.name as bee_farm_name, " +
			"hive_frame_types.name as hive_frame_type_name, hive_formats.name as hive_format_name, " +
			"hive_formats.size as hive_format_size").
		Where("hives.user_id = ? and hives.coord_x is null " +
			"and hives.coord_y is null and hives.deleted_at is null", id).
		Scan(&entities).Error

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

var DeleteHive = func(w http.ResponseWriter, r *http.Request) {
	var model models.Hive

	params := mux.Vars(r)
	id := params["id"]
	userID := u.GetUserIDFromRequest(r)

	db := db.GetDB()
	err := db.First(&model, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	if model.UserID != userID {
		u.HandleForbidden(w, errors.New("you are not allowed to do that"))
		return
	}

	err = db.Delete(&model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}