package ControllerLK

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
	"time"
)

type PollenHarvestEditModel struct {
	BeeFarmID     uint    `json:"bee_farm_id"`
	Amount        float64 `json:"amount"`
	Date          string  `json:"date"`
}

// TODO: add order by on different columns
// TODO: add pagination
var GetUsersPollenHarvests = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.PollenHarvest
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Preload("BeeFarm").
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

var CreatePollenHarvest = func(w http.ResponseWriter, r *http.Request) {
	EditModel := &PollenHarvestEditModel{}

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

	Model := models.PollenHarvest{UserID: id, Amount: EditModel.Amount,
		BeeFarmID: EditModel.BeeFarmID, Date: &t}

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}
