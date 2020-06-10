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

type BeeFarmCreateModel struct {
	Name          string   `json:"name"`
	Location      *string  `json:"location"`
	BeeFarmTypeID uint     `json:"bee_farm_type_id"`
	BeeFarmSizeID uint     `json:"bee_farm_size_id"`
}

type BeeFarmEditModel struct {
	Name          string   `json:"name"`
	Location      *string  `json:"location"`
	BeeFarmTypeID uint     `json:"bee_farm_type_id"`
}

var GetUsersBeeFarms = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.BeeFarm
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Preload("BeeFarmType").Preload("BeeFarmSize").
		Where("user_id = ?", id).Find(&entities).Error

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
	err := db.Preload("Reminders").Preload("BeeFamilies").Preload("BeeFarmSize").
		Preload("BeeFarmType").Preload("Hives").Preload("BeeFamilies.Swarms").
		Preload("BeeFamilies.BeeFamilyStatus").Preload("BeeFamilies.Hive").
		Preload("Hives.HiveFrameType").Preload("Hives.HiveFormat").
		Where("id = ?", id).Find(&BeeFarm).Error

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

var CreateBeeFarm = func(w http.ResponseWriter, r *http.Request) {
	CreateModel := &BeeFarmCreateModel{}

	err := json.NewDecoder(r.Body).Decode(CreateModel)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	id := u.GetUserIDFromRequest(r)
	Model := models.BeeFarm{Location: CreateModel.Location,
		Name: CreateModel.Name, BeeFarmTypeID: CreateModel.BeeFarmTypeID,
		BeeFarmSizeID: CreateModel.BeeFarmSizeID}
	Model.UserID = id

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var DeleteBeeFarm = func(w http.ResponseWriter, r *http.Request) {
	var model models.BeeFarm

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

var EditBeeFarm = func(w http.ResponseWriter, r *http.Request) {
	var model models.BeeFarm

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

	EditModel := &BeeFarmEditModel{}
	err = json.NewDecoder(r.Body).Decode(EditModel)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	model.Location = EditModel.Location
	model.Name = EditModel.Name
	model.BeeFarmTypeID = EditModel.BeeFarmTypeID

	err = db.Model(&models.BeeFarm{}).Updates(model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

