package AdminController

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

var BeeFamilyStatusCreate = func(w http.ResponseWriter, r *http.Request) {
	BeeFamilyStatus := &models.BeeFamilyStatus{}
	err := json.NewDecoder(r.Body).Decode(BeeFamilyStatus)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(BeeFamilyStatus).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(BeeFamilyStatus)
		u.RespondJSON(w, res)
	}
}

var BeeFamilyStatusRetrieve = func(w http.ResponseWriter, r *http.Request) {
	BeeFamilyStatus := &models.BeeFamilyStatus{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeFamilyStatus, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(BeeFamilyStatus)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if BeeFamilyStatus.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var BeeFamilyStatusUpdate = func(w http.ResponseWriter, r *http.Request) {
	BeeFamilyStatus := &models.BeeFamilyStatus{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeFamilyStatus, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newBeeFamilyStatus := &models.BeeFamilyStatus{}
	err = json.NewDecoder(r.Body).Decode(newBeeFamilyStatus)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&BeeFamilyStatus).Updates(newBeeFamilyStatus).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFamilyStatusDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.BeeFamilyStatus{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFamilyStatusQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.BeeFamilyStatus
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
		db.Model(&models.BeeFamilyStatus{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
