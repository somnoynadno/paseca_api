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

var SubscriptionTypeCreate = func(w http.ResponseWriter, r *http.Request) {
	SubscriptionType := &models.SubscriptionType{}
	err := json.NewDecoder(r.Body).Decode(SubscriptionType)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(SubscriptionType).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(SubscriptionType)
		u.RespondJSON(w, res)
	}
}

var SubscriptionTypeRetrieve = func(w http.ResponseWriter, r *http.Request) {
	SubscriptionType := &models.SubscriptionType{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&SubscriptionType, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(SubscriptionType)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if SubscriptionType.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var SubscriptionTypeUpdate = func(w http.ResponseWriter, r *http.Request) {
	SubscriptionType := &models.SubscriptionType{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&SubscriptionType, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newSubscriptionType := &models.SubscriptionType{}
	err = json.NewDecoder(r.Body).Decode(newSubscriptionType)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&SubscriptionType).Updates(newSubscriptionType).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var SubscriptionTypeDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.SubscriptionType{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var SubscriptionTypeQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.SubscriptionType
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
		db.Model(&models.SubscriptionType{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
