package LKCustomCRUD

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
)

var GetBeeFarmSizes = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.BeeFarmSize
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Where("is_custom = false or (is_custom = true and creator_id = ?)", id).Find(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var DeleteBeeFarmSizeByID = func(w http.ResponseWriter, r *http.Request) {
	var entity models.BeeFarmSize

	params := mux.Vars(r)
	id := params["id"]

	userID := u.GetUserIDFromRequest(r)

	db := db.GetDB()
	err := db.First(&entity, id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			u.HandleNotFound(w)
		} else {
			u.HandleBadRequest(w, err)
		}
		return
	}

	if entity.CreatorID != &userID {
		u.HandleForbidden(w, errors.New("you are not allowed to do that"))
		return
	}

	res, err := json.Marshal(entity)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}
