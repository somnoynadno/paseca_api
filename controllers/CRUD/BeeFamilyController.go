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

var BeeFamilyCreate = func(w http.ResponseWriter, r *http.Request) {
	BeeFamily := &models.BeeFamily{}
	err := json.NewDecoder(r.Body).Decode(BeeFamily)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(BeeFamily).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(BeeFamily)
		u.RespondJSON(w, res)
	}
}

var BeeFamilyRetrieve = func(w http.ResponseWriter, r *http.Request) {
	BeeFamily := &models.BeeFamily{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeFamily, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(BeeFamily)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if BeeFamily.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var BeeFamilyUpdate = func(w http.ResponseWriter, r *http.Request) {
	BeeFamily := &models.BeeFamily{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeFamily, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newBeeFamily := &models.BeeFamily{}
	err = json.NewDecoder(r.Body).Decode(newBeeFamily)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&BeeFamily).Updates(newBeeFamily).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFamilyDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.BeeFamily{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFamilyQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.BeeFamily
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
		db.Model(&models.BeeFamily{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
