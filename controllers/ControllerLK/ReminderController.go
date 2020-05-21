package ControllerLK

import (
	"encoding/json"
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

var CreateReminder= func(w http.ResponseWriter, r *http.Request) {
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

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}