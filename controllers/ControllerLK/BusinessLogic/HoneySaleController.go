package BusinessLogic

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
	"time"
)

type HoneySaleEditModel struct {
	HoneyTypeID   uint    `json:"honey_type_id"`
	BeeFarmID     uint    `json:"bee_farm_id"`
	Amount        float64 `json:"amount"`
	TotalPrice    float64 `json:"total_price"`
	Date          string  `json:"date"`
}

// TODO: add order by on different columns
// TODO: add pagination
var GetUsersHoneySales = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.HoneySale
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Preload("HoneyType").Preload("BeeFarm").
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

var CreateHoneySale = func(w http.ResponseWriter, r *http.Request) {
	EditModel := &HoneySaleEditModel{}

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

	Model := models.HoneySale{UserID: id, Amount: EditModel.Amount,
		TotalPrice: EditModel.TotalPrice, HoneyTypeID: EditModel.HoneyTypeID,
		BeeFarmID: EditModel.BeeFarmID, Date: &t}

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}