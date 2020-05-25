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
	"time"
)

type ControlHarvestEditModel struct {
	BeeFamilyID   uint    `json:"bee_family_id"`
	Amount        float64 `json:"amount"`
	Date          string  `json:"date"`
}

// TODO: add order by on different columns
// TODO: add pagination
var GetUsersControlHarvests = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.ControlHarvest
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Preload("BeeFamily").
		Where("user_id = ?", id).Order("date desc").Find(&entities).Error

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

var CreateControlHarvest = func(w http.ResponseWriter, r *http.Request) {
	EditModel := &ControlHarvestEditModel{}

	err := json.NewDecoder(r.Body).Decode(EditModel)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	id := u.GetUserIDFromRequest(r)
	t, err := time.Parse("2006-01-02", EditModel.Date)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	Model := models.ControlHarvest{Amount: EditModel.Amount,
		BeeFamilyID: EditModel.BeeFamilyID, Date: &t}
	Model.UserID = id

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var DeleteControlHarvest = func(w http.ResponseWriter, r *http.Request) {
	var model models.ControlHarvest

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
