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

var FamilyDiseaseCreate = func(w http.ResponseWriter, r *http.Request) {
	FamilyDisease := &models.FamilyDisease{}
	err := json.NewDecoder(r.Body).Decode(FamilyDisease)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(FamilyDisease).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(FamilyDisease)
		u.RespondJSON(w, res)
	}
}

var FamilyDiseaseRetrieve = func(w http.ResponseWriter, r *http.Request) {
	FamilyDisease := &models.FamilyDisease{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("BeeFamily").Preload("BeeDisease").First(&FamilyDisease, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(FamilyDisease)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if FamilyDisease.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var FamilyDiseaseUpdate = func(w http.ResponseWriter, r *http.Request) {
	FamilyDisease := &models.FamilyDisease{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&FamilyDisease, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newFamilyDisease := &models.FamilyDisease{}
	err = json.NewDecoder(r.Body).Decode(newFamilyDisease)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	// do NOT update recursively
	newFamilyDisease.BeeFamily = models.BeeFamily{}
	newFamilyDisease.BeeDisease = models.BeeDisease{}

	err = db.Model(&FamilyDisease).Updates(newFamilyDisease).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var FamilyDiseaseDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.FamilyDisease{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var FamilyDiseaseQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.FamilyDisease
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
	err := db.Preload("BeeFamily").Preload("BeeDisease").
		Order(fmt.Sprintf("%s %s", sort, order)).Offset(start).Limit(end + start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		db.Model(&models.FamilyDisease{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
