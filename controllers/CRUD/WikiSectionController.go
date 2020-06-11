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

var WikiSectionCreate = func(w http.ResponseWriter, r *http.Request) {
	WikiSection := &models.WikiSection{}
	err := json.NewDecoder(r.Body).Decode(WikiSection)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(WikiSection).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(WikiSection)
		u.RespondJSON(w, res)
	}
}

var WikiSectionRetrieve = func(w http.ResponseWriter, r *http.Request) {
	WikiSection := &models.WikiSection{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&WikiSection, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(WikiSection)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if WikiSection.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var WikiSectionUpdate = func(w http.ResponseWriter, r *http.Request) {
	WikiSection := &models.WikiSection{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&WikiSection, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newWikiSection := &models.WikiSection{}
	err = json.NewDecoder(r.Body).Decode(newWikiSection)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	err = db.Model(&WikiSection).Updates(newWikiSection).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var WikiSectionDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.WikiSection{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var WikiSectionQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.WikiSection
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
		db.Model(&models.WikiSection{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
