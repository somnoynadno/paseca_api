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

var HiveCreate = func(w http.ResponseWriter, r *http.Request) {
	Hive := &models.Hive{}
	err := json.NewDecoder(r.Body).Decode(Hive)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(Hive).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(Hive)
		u.RespondJSON(w, res)
	}
}

var HiveRetrieve = func(w http.ResponseWriter, r *http.Request) {
	Hive := &models.Hive{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("BeeFamily").Preload("HiveFormat").Preload("HiveFrameType").
		First(&Hive, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(Hive)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if Hive.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var HiveUpdate = func(w http.ResponseWriter, r *http.Request) {
	Hive := &models.Hive{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&Hive, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newHive := &models.Hive{}
	err = json.NewDecoder(r.Body).Decode(newHive)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	// do NOT update recursively
	newHive.BeeFamily = new(models.BeeFamily)
	newHive.HiveFormat = models.HiveFormat{}
	newHive.HiveFrameType = models.HiveFrameType{}

	err = db.Model(&Hive).Updates(newHive).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HiveDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.Hive{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var HiveQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.Hive
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
	err := db.Preload("BeeFamily").Preload("HiveFormat").Preload("HiveFrameType").
		Order(fmt.Sprintf("%s %s", sort, order)).Offset(start).Limit(end + start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		db.Model(&models.Hive{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
