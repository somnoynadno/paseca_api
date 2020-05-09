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

var HoneyHarvestCreate = func(w http.ResponseWriter, r *http.Request) {
	HoneyHarvest := &models.HoneyHarvest{}
	err := json.NewDecoder(r.Body).Decode(HoneyHarvest)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(HoneyHarvest).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(HoneyHarvest)
		u.RespondJSON(w, res)
	}
}

var HoneyHarvestRetrieve = func(w http.ResponseWriter, r *http.Request) {
	HoneyHarvest := &models.HoneyHarvest{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&HoneyHarvest, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(HoneyHarvest)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if HoneyHarvest.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var HoneyHarvestUpdate = func(w http.ResponseWriter, r *http.Request) {
	HoneyHarvest := &models.HoneyHarvest{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&HoneyHarvest, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newHoneyHarvest := &models.HoneyHarvest{}
	err = json.NewDecoder(r.Body).Decode(newHoneyHarvest)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&HoneyHarvest).Updates(newHoneyHarvest).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HoneyHarvestDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.HoneyHarvest{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HoneyHarvestQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.HoneyHarvest
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
		db.Model(&models.HoneyHarvest{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
