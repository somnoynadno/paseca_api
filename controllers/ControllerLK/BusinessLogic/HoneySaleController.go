package BusinessLogic

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
	"strconv"
	"time"
)

type HoneySaleEditModel struct {
	HoneyTypeID   uint    `json:"honey_type_id"`
	BeeFarmID     uint    `json:"bee_farm_id"`
	Amount        float64 `json:"amount"`
	TotalPrice    float64 `json:"total_price"`
	Date          string  `json:"date"`
}

var GetUsersHoneySales = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.HoneySale
	id := r.Context().Value("context").(u.Values).Get("user_id")

	order := r.FormValue("_order")
	sort := r.FormValue("_sort")
	end, err1 := strconv.Atoi(r.FormValue("_end"))
	start, err2 := strconv.Atoi(r.FormValue("_start"))

	if err1 != nil || err2 != nil {
		u.HandleBadRequest(w, errors.New("bad _start or _end parameter value"))
		return
	}
	u.CheckOrderAndSortParams(&order, &sort)

	db := db.GetDB()
	err := db.Preload("HoneyType").Preload("BeeFarm").Where("user_id = ?", id).
		Order(fmt.Sprintf("%s %s", sort, order)).Offset(start).Limit(end - start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		var count string
		db.Model(&models.HoneySale{}).Where("user_id = ?", id).Count(&count)
		u.SetTotalCountHeader(w, count)
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

	Model := models.HoneySale{Amount: EditModel.Amount,
		TotalPrice: EditModel.TotalPrice, HoneyTypeID: EditModel.HoneyTypeID,
		BeeFarmID: EditModel.BeeFarmID, Date: &t}
	Model.UserID = id

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var DeleteHoneySale = func(w http.ResponseWriter, r *http.Request) {
	var model models.HoneySale

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
