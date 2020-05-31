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

var SwarmStatusCreate = func(w http.ResponseWriter, r *http.Request) {
	SwarmStatus := &models.SwarmStatus{}
	err := json.NewDecoder(r.Body).Decode(SwarmStatus)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(SwarmStatus).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(SwarmStatus)
		u.RespondJSON(w, res)
	}
}

var SwarmStatusRetrieve = func(w http.ResponseWriter, r *http.Request) {
	SwarmStatus := &models.SwarmStatus{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&SwarmStatus, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(SwarmStatus)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if SwarmStatus.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var SwarmStatusUpdate = func(w http.ResponseWriter, r *http.Request) {
	SwarmStatus := &models.SwarmStatus{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&SwarmStatus, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newSwarmStatus := &models.SwarmStatus{}
	err = json.NewDecoder(r.Body).Decode(newSwarmStatus)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&SwarmStatus).Updates(newSwarmStatus).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var SwarmStatusDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.SwarmStatus{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var SwarmStatusQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.SwarmStatus
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
		db.Model(&models.SwarmStatus{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
