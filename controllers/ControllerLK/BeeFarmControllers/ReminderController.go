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

type ReminderCreateModel struct {
	Title     string  `json:"title"`
	Text      string  `json:"text"`
	Date      string  `json:"date"`
	BeeFarmID uint    `json:"bee_farm_id"`
}

var CreateReminder = func(w http.ResponseWriter, r *http.Request) {
	CreateModel := &ReminderCreateModel{}

	err := json.NewDecoder(r.Body).Decode(CreateModel)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	date, _ := time.Parse("2006-01-02", CreateModel.Date)

	Model := models.Reminder{BeeFarmID: CreateModel.BeeFarmID,
		Title: CreateModel.Title, Text: CreateModel.Text, Date: date,
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

var CheckReminder = func(w http.ResponseWriter, r *http.Request) {
	var model models.Reminder

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

	model.IsChecked = !model.IsChecked
	err = db.Save(&model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var DeleteReminder = func(w http.ResponseWriter, r *http.Request) {
	var model models.Reminder

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

var GetRemindersByBeeFarmID = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.Reminder

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Where("bee_farm_id = ?", id).
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