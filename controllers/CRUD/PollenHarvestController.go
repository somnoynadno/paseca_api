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

var PollenHarvestCreate = func(w http.ResponseWriter, r *http.Request) {
	PollenHarvest := &models.PollenHarvest{}
	err := json.NewDecoder(r.Body).Decode(PollenHarvest)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(PollenHarvest).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(PollenHarvest)
		u.RespondJSON(w, res)
	}
}

var PollenHarvestRetrieve = func(w http.ResponseWriter, r *http.Request) {
	PollenHarvest := &models.PollenHarvest{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("BeeFarm").First(&PollenHarvest, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(PollenHarvest)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if PollenHarvest.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var PollenHarvestUpdate = func(w http.ResponseWriter, r *http.Request) {
	PollenHarvest := &models.PollenHarvest{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&PollenHarvest, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newPollenHarvest := &models.PollenHarvest{}
	err = json.NewDecoder(r.Body).Decode(newPollenHarvest)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	// do NOT update recursively
	newPollenHarvest.BeeFarm = models.BeeFarm{}

	err = db.Model(&PollenHarvest).Updates(newPollenHarvest).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var PollenHarvestDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.PollenHarvest{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var PollenHarvestQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.PollenHarvest
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
	err := db.Preload("BeeFarm").
		Order(fmt.Sprintf("%s %s", sort, order)).Offset(start).Limit(end + start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		db.Model(&models.PollenHarvest{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
