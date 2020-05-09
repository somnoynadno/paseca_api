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

var BeeDiseaseCreate = func(w http.ResponseWriter, r *http.Request) {
	BeeDisease := &models.BeeDisease{}
	err := json.NewDecoder(r.Body).Decode(BeeDisease)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(BeeDisease).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(BeeDisease)
		u.RespondJSON(w, res)
	}
}

var BeeDiseaseRetrieve = func(w http.ResponseWriter, r *http.Request) {
	BeeDisease := &models.BeeDisease{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeDisease, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(BeeDisease)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if BeeDisease.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var BeeDiseaseUpdate = func(w http.ResponseWriter, r *http.Request) {
	BeeDisease := &models.BeeDisease{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeDisease, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newBeeDisease := &models.BeeDisease{}
	err = json.NewDecoder(r.Body).Decode(newBeeDisease)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&BeeDisease).Updates(newBeeDisease).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeDiseaseDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.BeeDisease{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeDiseaseQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.BeeDisease
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
		db.Model(&models.BeeDisease{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
