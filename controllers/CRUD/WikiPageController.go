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

var WikiPageCreate = func(w http.ResponseWriter, r *http.Request) {
	WikiPage := &models.WikiPage{}
	err := json.NewDecoder(r.Body).Decode(WikiPage)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	db := db.GetDB()
	err = db.Create(WikiPage).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		res, _ := json.Marshal(WikiPage)
		u.RespondJSON(w, res)
	}
}

var WikiPageRetrieve = func(w http.ResponseWriter, r *http.Request) {
	WikiPage := &models.WikiPage{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("WikiSection").First(&WikiPage, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	res, err := json.Marshal(WikiPage)
	if err != nil {
		u.HandleBadRequest(w, err)
	} else if WikiPage.ID == 0 {
		u.HandleNotFound(w)
	} else {
		u.RespondJSON(w, res)
	}
}

var WikiPageUpdate = func(w http.ResponseWriter, r *http.Request) {
	WikiPage := &models.WikiPage{}

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.First(&WikiPage, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	newWikiPage := &models.WikiPage{}
	err = json.NewDecoder(r.Body).Decode(newWikiPage)

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	// do NOT update recursively
	newWikiPage.WikiSection = WikiPage.WikiSection

	err = db.Model(&WikiPage).Updates(newWikiPage).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var WikiPageDelete = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Delete(&models.WikiPage{}, id).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}

var WikiPageQuery = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.WikiPage
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
	err := db.Preload("WikiSection").Order(fmt.Sprintf("%s %s", sort, order)).
		Offset(start).Limit(end + start).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		db.Model(&models.WikiPage{}).Count(&count)
		u.SetTotalCountHeader(w, count)
		u.RespondJSON(w, res)
	}
}
