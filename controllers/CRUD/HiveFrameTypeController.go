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

var HiveFrameTypeCreate = func(w http.ResponseWriter, r *http.Request) {
	HiveFrameType := &models.HiveFrameType{}
	err := json.NewDecoder(r.Body).Decode(HiveFrameType)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(HiveFrameType).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(HiveFrameType)
		u.RespondJSON(w, res)
	}
}

var HiveFrameTypeRetrieve = func(w http.ResponseWriter, r *http.Request) {
	HiveFrameType := &models.HiveFrameType{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&HiveFrameType, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(HiveFrameType)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if HiveFrameType.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var HiveFrameTypeUpdate = func(w http.ResponseWriter, r *http.Request) {
	HiveFrameType := &models.HiveFrameType{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&HiveFrameType, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newHiveFrameType := &models.HiveFrameType{}
	err = json.NewDecoder(r.Body).Decode(newHiveFrameType)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&HiveFrameType).Updates(newHiveFrameType).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HiveFrameTypeDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.HiveFrameType{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HiveFrameTypeQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.HiveFrameType
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
		db.Model(&models.HiveFrameType{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
