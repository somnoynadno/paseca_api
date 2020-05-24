package BusinessLogic

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
	"time"
)

type HoneyHarvestEditModel struct {
	HoneyTypeID   uint    `json:"honey_type_id"`
	BeeFamilyID   uint    `json:"bee_family_id"`
	Amount        float64 `json:"amount"`
	Date          string  `json:"date"`
}

// TODO: add order by on different columns
// TODO: add pagination
var GetUsersHoneyHarvests = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.HoneyHarvest
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Preload("HoneyType").Preload("BeeFamily").
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

var CreateHoneyHarvest = func(w http.ResponseWriter, r *http.Request) {
	EditModel := &HoneyHarvestEditModel{}

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

	// TODO: predict total price based on honey type and amount
	Model := models.HoneyHarvest{UserID: id, Amount: EditModel.Amount,
		TotalPrice: 0, HoneyTypeID: EditModel.HoneyTypeID,
		BeeFamilyID: EditModel.BeeFamilyID, Date: &t}

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}
