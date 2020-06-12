package AdminController

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
)

var SubscriptionStatusCreate = func(w http.ResponseWriter, r *http.Request) {
	SubscriptionStatus := &models.SubscriptionStatus{}
	err := json.NewDecoder(r.Body).Decode(SubscriptionStatus)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(SubscriptionStatus).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(SubscriptionStatus)
		u.RespondJSON(w, res)
	}
}

var SubscriptionStatusRetrieve = func(w http.ResponseWriter, r *http.Request) {
	SubscriptionStatus := &models.SubscriptionStatus{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&SubscriptionStatus, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(SubscriptionStatus)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if SubscriptionStatus.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var SubscriptionStatusUpdate = func(w http.ResponseWriter, r *http.Request) {
	SubscriptionStatus := &models.SubscriptionStatus{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&SubscriptionStatus, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newSubscriptionStatus := &models.SubscriptionStatus{}
	err = json.NewDecoder(r.Body).Decode(newSubscriptionStatus)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&SubscriptionStatus).Updates(newSubscriptionStatus).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var SubscriptionStatusDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.SubscriptionStatus{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var SubscriptionStatusQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.SubscriptionStatus
	var count string

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
	err := db.Order(fmt.Sprintf("%s %s", sort, order)).Offset(start).Limit(end + start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		db.Model(&models.SubscriptionStatus{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
