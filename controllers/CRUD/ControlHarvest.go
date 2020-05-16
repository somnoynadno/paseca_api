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

var ControlHarvestCreate = func(w http.ResponseWriter, r *http.Request) {
	ControlHarvest := &models.ControlHarvest{}
	err := json.NewDecoder(r.Body).Decode(ControlHarvest)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(ControlHarvest).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(ControlHarvest)
		u.RespondJSON(w, res)
	}
}

var ControlHarvestRetrieve = func(w http.ResponseWriter, r *http.Request) {
	ControlHarvest := &models.ControlHarvest{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("BeeFamily").First(&ControlHarvest, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(ControlHarvest)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if ControlHarvest.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var ControlHarvestUpdate = func(w http.ResponseWriter, r *http.Request) {
	ControlHarvest := &models.ControlHarvest{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&ControlHarvest, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newControlHarvest := &models.ControlHarvest{}
	err = json.NewDecoder(r.Body).Decode(newControlHarvest)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	// do NOT update recursively
	newControlHarvest.BeeFamily = models.BeeFamily{}

	err = db.Model(&ControlHarvest).Updates(newControlHarvest).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var ControlHarvestDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.ControlHarvest{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var ControlHarvestQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.ControlHarvest
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
	err := db.Preload("BeeFamily").
		Order(fmt.Sprintf("%s %s", sort, order)).Offset(start).Limit(end + start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		db.Model(&models.ControlHarvest{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
