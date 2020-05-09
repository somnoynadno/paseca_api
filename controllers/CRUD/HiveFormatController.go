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

var HiveFormatCreate = func(w http.ResponseWriter, r *http.Request) {
	HiveFormat := &models.HiveFormat{}
	err := json.NewDecoder(r.Body).Decode(HiveFormat)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(HiveFormat).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(HiveFormat)
		u.RespondJSON(w, res)
	}
}

var HiveFormatRetrieve = func(w http.ResponseWriter, r *http.Request) {
	HiveFormat := &models.HiveFormat{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&HiveFormat, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(HiveFormat)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if HiveFormat.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var HiveFormatUpdate = func(w http.ResponseWriter, r *http.Request) {
	HiveFormat := &models.HiveFormat{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&HiveFormat, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newHiveFormat := &models.HiveFormat{}
	err = json.NewDecoder(r.Body).Decode(newHiveFormat)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&HiveFormat).Updates(newHiveFormat).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HiveFormatDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.HiveFormat{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HiveFormatQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.HiveFormat
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
		db.Model(&models.HiveFormat{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
