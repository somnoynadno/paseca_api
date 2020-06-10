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

type FamilyDiseaseCreateModel struct {
	BeeFamilyID  uint `json:"bee_family_id"`
	BeeDiseaseID uint `json:"bee_disease_id"`
}

var CreateFamilyDisease = func(w http.ResponseWriter, r *http.Request) {
	CreateModel := &FamilyDiseaseCreateModel{}

	err := json.NewDecoder(r.Body).Decode(CreateModel)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	Model := models.FamilyDisease{
		BeeFamilyID: CreateModel.BeeFamilyID, BeeDiseaseID: CreateModel.BeeDiseaseID,
	}
	Model.UserID = u.GetUserIDFromRequest(r)

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var GetFamilyDiseasesByBeeFarmID = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.FamilyDisease

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("BeeDisease").
		Preload("BeeFamily", "bee_farm_id = ?", id).
		Find(&entities).Error

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

var DeleteFamilyDisease = func(w http.ResponseWriter, r *http.Request) {
	var model models.FamilyDisease

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