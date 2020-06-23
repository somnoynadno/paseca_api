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
	"time"
)

type UserCreateModel struct {
	Email                string             `json:"email"`
	Name                 string             `json:"name"`
	Surname              string             `json:"surname"`
	Password             string             `json:"password"`
	IsAdmin              bool               `json:"is_admin"`
	SubscriptionEnd      *time.Time         `json:"subscription_end"`
	SubscriptionStatusID uint               `json:"subscription_status_id"`
	SubscriptionTypeID   uint               `json:"subscription_type_id"`
}

var UserCreate = func(w http.ResponseWriter, r *http.Request) {
	User := &UserCreateModel{}
	err := json.NewDecoder(r.Body).Decode(User)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	hash, err := u.HashPassword(User.Password)
	if err != nil {
		u.HandleInternalError(w, err)
		return
	}
	User.Password = hash

	db := db.GetDB()
	err = db.Create(User).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(User)
		u.RespondJSON(w, res)
	}
}

var UserRetrieve = func(w http.ResponseWriter, r *http.Request) {
	User := &models.User{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("SubscriptionStatus").Preload("SubscriptionType").
		Preload("BeeFarms").First(&User, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(User)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if User.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var UserUpdate = func(w http.ResponseWriter, r *http.Request) {
	User := &models.User{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&User, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newUser := &models.User{}
	err = json.NewDecoder(r.Body).Decode(newUser)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	// do NOT update recursively
	newUser.SubscriptionType = models.SubscriptionType{}
	newUser.SubscriptionStatus = models.SubscriptionStatus{}

	db.Model(&User).Update("is_admin", newUser.IsAdmin)
	err = db.Model(&User).Updates(newUser).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var UserDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.User{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var UserQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.User
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
	err := db.Preload("SubscriptionStatus").Preload("SubscriptionType").
		Order(fmt.Sprintf("%s %s", sort, order)).Offset(start).Limit(end + start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		db.Model(&models.User{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
