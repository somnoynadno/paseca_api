package BeeFarmControllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
)

var GetSwarmsByBeeFarmID = func(w http.ResponseWriter, r *http.Request) {
	var entities []models.Swarm

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("SwarmStatus").
		Preload("BeeFamily", "bee_farm_id = ?", id).
		Find(&entities).Error

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
