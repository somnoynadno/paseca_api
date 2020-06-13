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

var BeeFarmSizeCreate = func(w http.ResponseWriter, r *http.Request) {
	BeeFarmSize := &models.BeeFarmSize{}
	err := json.NewDecoder(r.Body).Decode(BeeFarmSize)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(BeeFarmSize).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(BeeFarmSize)
		u.RespondJSON(w, res)
	}
}

var BeeFarmSizeRetrieve = func(w http.ResponseWriter, r *http.Request) {
	BeeFarmSize := &models.BeeFarmSize{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeFarmSize, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(BeeFarmSize)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if BeeFarmSize.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var BeeFarmSizeUpdate = func(w http.ResponseWriter, r *http.Request) {
	BeeFarmSize := &models.BeeFarmSize{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeFarmSize, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newBeeFarmSize := &models.BeeFarmSize{}
	err = json.NewDecoder(r.Body).Decode(newBeeFarmSize)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&BeeFarmSize).Updates(newBeeFarmSize).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFarmSizeDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.BeeFarmSize{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFarmSizeQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.BeeFarmSize
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
		db.Model(&models.BeeFarmSize{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
