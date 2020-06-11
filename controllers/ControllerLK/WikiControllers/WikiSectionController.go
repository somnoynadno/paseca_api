package WikiControllers

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
)

var GetWikiSections = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.WikiSection

	db := db.GetDB()
	err := db.Find(&entities).Error

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
