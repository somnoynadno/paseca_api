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
	"time"
)

type SwarmCreateModel struct {
	Date          *string   `json:"date"`
	Order         *int      `json:"order"`
	BeeFamilyID   uint      `json:"bee_family_id"`
	SwarmStatusID uint      `json:"swarm_status_id"`
}

var CreateSwarm = func(w http.ResponseWriter, r *http.Request) {
	CreateModel := &SwarmCreateModel{}

	err := json.NewDecoder(r.Body).Decode(CreateModel)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	date, _ := time.Parse("2006-01-02", *CreateModel.Date)

	Model := models.Swarm{Date: &date, Order: CreateModel.Order,
		BeeFamilyID: CreateModel.BeeFamilyID, SwarmStatusID: CreateModel.SwarmStatusID,
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

var GetSwarmsByBeeFarmID = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.Swarm

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Joins("join bee_families on bee_families.id = swarms.bee_family_id").
		Where("bee_families.bee_farm_id = ?", id).Preload("SwarmStatus").
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

var DeleteSwarm = func(w http.ResponseWriter, r *http.Request) {
	var model models.Swarm

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
