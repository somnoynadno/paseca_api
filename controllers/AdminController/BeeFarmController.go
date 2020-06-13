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

var BeeFarmCreate = func(w http.ResponseWriter, r *http.Request) {
	BeeFarm := &models.BeeFarm{}
	err := json.NewDecoder(r.Body).Decode(BeeFarm)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(BeeFarm).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(BeeFarm)
		u.RespondJSON(w, res)
	}
}

var BeeFarmRetrieve = func(w http.ResponseWriter, r *http.Request) {
	BeeFarm := &models.BeeFarm{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("User").Preload("BeeFarmType").Preload("BeeFarmSize").
		Preload("Reminders").Preload("BeeFamilies").Preload("HoneySales").Preload("Hives").
		First(&BeeFarm, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(BeeFarm)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if BeeFarm.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var BeeFarmUpdate = func(w http.ResponseWriter, r *http.Request) {
	BeeFarm := &models.BeeFarm{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&BeeFarm, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newBeeFarm := &models.BeeFarm{}
	err = json.NewDecoder(r.Body).Decode(newBeeFarm)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	// do NOT update recursively
	newBeeFarm.User = BeeFarm.User
	newBeeFarm.BeeFarmType = models.BeeFarmType{}
	newBeeFarm.BeeFarmSize = models.BeeFarmSize{}

	err = db.Model(&BeeFarm).Updates(newBeeFarm).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFarmDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.BeeFarm{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var BeeFarmQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.BeeFarm
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
	err := db.Preload("User").Preload("BeeFarmType").Preload("BeeFarmSize").
		Order(fmt.Sprintf("%s %s", sort, order)).Offset(start).Limit(end + start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		db.Model(&models.BeeFarm{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
