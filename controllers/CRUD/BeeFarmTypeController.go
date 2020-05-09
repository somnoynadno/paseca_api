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

var BeeFarmTypeCreate = func(w http.ResponseWriter, r *http.Request) {
	BeeFarmType := &models.BeeFarmType{}
	err := json.NewDecoder(r.Body).Decode(BeeFarmType)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(BeeFarmType).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(BeeFarmType)
		u.RespondJSON(w, res)
	}
}

var BeeFarmTypeRetrieve = func(w http.ResponseWriter, r *http.Request) {
	BeeFarmType := &models.BeeFarmType{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeFarmType, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(BeeFarmType)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if BeeFarmType.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var BeeFarmTypeUpdate = func(w http.ResponseWriter, r *http.Request) {
	BeeFarmType := &models.BeeFarmType{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeFarmType, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newBeeFarmType := &models.BeeFarmType{}
	err = json.NewDecoder(r.Body).Decode(newBeeFarmType)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&BeeFarmType).Updates(newBeeFarmType).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFarmTypeDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.BeeFarmType{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFarmTypeQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.BeeFarmType
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
		db.Model(&models.BeeFarmType{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
