package CRUD

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

var HoneySaleCreate = func(w http.ResponseWriter, r *http.Request) {
	HoneySale := &models.HoneySale{}
	err := json.NewDecoder(r.Body).Decode(HoneySale)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(HoneySale).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(HoneySale)
		u.RespondJSON(w, res)
	}
}

var HoneySaleRetrieve = func(w http.ResponseWriter, r *http.Request) {
	HoneySale := &models.HoneySale{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&HoneySale, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(HoneySale)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if HoneySale.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var HoneySaleUpdate = func(w http.ResponseWriter, r *http.Request) {
	HoneySale := &models.HoneySale{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&HoneySale, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newHoneySale := &models.HoneySale{}
	err = json.NewDecoder(r.Body).Decode(newHoneySale)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&HoneySale).Updates(newHoneySale).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HoneySaleDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.HoneySale{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HoneySaleQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.HoneySale
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
		db.Model(&models.HoneySale{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
