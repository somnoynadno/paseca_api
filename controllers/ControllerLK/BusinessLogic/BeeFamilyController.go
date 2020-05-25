package BusinessLogic

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
	. "time"
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

var GetUsersBeeFamiliesWithoutHives = func(w http.ResponseWriter, r *http.Request) {
	var entities []BeeFamilyShort
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Table("bee_families").
		Joins("join bee_farms on bee_farms.id = bee_families.bee_farm_id").
		Joins("join users on users.id = bee_farms.user_id").
		Select("bee_families.id, bee_families.name, bee_farms.name as bee_farm_name").
		Where("users.id = ? and bee_families.hive_id is null", id).Scan(&entities).Error

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

type BeeFamilyCreateModel struct {
	Name               string   `json:"name"`
	QueenBeeBornDate   *string  `json:"queen_bee_born_date"`
	LastInspectionDate *string  `json:"last_inspection_date"`
	BeeFarmID          uint     `json:"bee_farm_id"`
	BeeBreedID         uint     `json:"bee_breed_id"`
	BeeFamilyStatusID  uint     `json:"bee_family_status_id"`
	Parent1ID          *uint    `json:"parent1_id"`
	Parent2ID          *uint    `json:"parent2_id"`
	IsControl          bool     `json:"is_control"`
}

var CreateBeeFamily = func(w http.ResponseWriter, r *http.Request) {
	CreateModel := &BeeFamilyCreateModel{}

	err := json.NewDecoder(r.Body).Decode(CreateModel)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	var queenBeeBornDate *Time
	var lastInspectionDate *Time

	if CreateModel.QueenBeeBornDate != nil {
		t, _ := Parse("2006-01-02", *CreateModel.QueenBeeBornDate)
		queenBeeBornDate = &t
	}
	if CreateModel.LastInspectionDate != nil {
		t, _ := Parse("2006-01-02", *CreateModel.LastInspectionDate)
		lastInspectionDate = &t
	}

	Model := models.BeeFamily{BeeFarmID: CreateModel.BeeFarmID,
		Name: CreateModel.Name, BeeBreedID: CreateModel.BeeBreedID,
		QueenBeeBornDate: queenBeeBornDate, LastInspectionDate: lastInspectionDate,
		Parent1ID: CreateModel.Parent1ID, Parent2ID: CreateModel.Parent2ID,
		IsControl: CreateModel.IsControl, BeeFamilyStatusID: CreateModel.BeeFamilyStatusID,
	}

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var GetBeeFamilyByID = func(w http.ResponseWriter, r *http.Request) {
	var BeeFamily models.BeeFamily

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("BeeBreed").Preload("BeeFamilyStatus").Preload("Hive").
		Preload("BeeDiseases").Preload("HoneyHarvests").Preload("ControlHarvests").
		Preload("Parent1").Preload("Parent2").
		Where("id = ?", id).Find(&BeeFamily).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(BeeFamily)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var DeleteBeeFamily = func(w http.ResponseWriter, r *http.Request) {
	var model models.BeeFamily

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