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

var SwarmCreate = func(w http.ResponseWriter, r *http.Request) {
	Swarm := &models.Swarm{}
	err := json.NewDecoder(r.Body).Decode(Swarm)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(Swarm).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(Swarm)
		u.RespondJSON(w, res)
	}
}

var SwarmRetrieve = func(w http.ResponseWriter, r *http.Request) {
	Swarm := &models.Swarm{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("SwarmStatus").Preload("BeeFamily").First(&Swarm, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(Swarm)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if Swarm.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var SwarmUpdate = func(w http.ResponseWriter, r *http.Request) {
	Swarm := &models.Swarm{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&Swarm, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newSwarm := &models.Swarm{}
	err = json.NewDecoder(r.Body).Decode(newSwarm)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	newSwarm.SwarmStatus = Swarm.SwarmStatus
	newSwarm.BeeFamily = Swarm.BeeFamily

	err = db.Model(&Swarm).Updates(newSwarm).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var SwarmDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.Swarm{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var SwarmQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.Swarm
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
	err := db.Preload("SwarmStatus").Preload("BeeFamily").
		Order(fmt.Sprintf("%s %s", sort, order)).Offset(start).Limit(end + start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		db.Model(&models.Swarm{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
