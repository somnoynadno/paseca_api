package BusinessLogic

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"paseca/db"
	"paseca/models"
	u "paseca/utils"
)

type BeeFarm struct {
	models.BaseModel
	Name          string             `json:"name"`
	Location      *string            `json:"location"`
	BeeFarmTypeID uint               `json:"bee_farm_type_id"`
	BeeFarmType   models.BeeFarmType `json:"bee_farm_type"`
	BeeFarmSizeID uint               `json:"bee_farm_size_id"`
	BeeFarmSize   models.BeeFarmSize `json:"bee_farm_size"`
}

type BeeFarmEditModel struct {
	Name          string   `json:"name"`
	Location      *string  `json:"location"`
	BeeFarmTypeID uint     `json:"bee_farm_type_id"`
	BeeFarmSizeID uint     `json:"bee_farm_size_id"`
}

var GetUsersBeeFarms = func(w http.ResponseWriter, r *http.Request) {
	var entities []BeeFarm
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Preload("BeeFarmType").Preload("BeeFarmSize").
		Where("user_id = ?", id).Find(&entities).Error

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

var GetBeeFarmByID = func(w http.ResponseWriter, r *http.Request) {
	var BeeFarm models.BeeFarm

	params := mux.Vars(r)
	id := params["id"]

	db := db.GetDB()
	err := db.Preload("Reminders").Preload("BeeFamilies").Preload("BeeFarmSize").
		Preload("BeeFarmType").Preload("Hives").
		Preload("BeeFamilies.BeeFamilyStatus").Preload("BeeFamilies.Hive").
		Preload("Hives.HiveFrameType").Preload("Hives.HiveFormat").
		Where("id = ?", id).Find(&BeeFarm).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	res, err := json.Marshal(BeeFarm)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}

var CreateBeeFarm = func(w http.ResponseWriter, r *http.Request) {
	EditModel := &BeeFarmEditModel{}

	err := json.NewDecoder(r.Body).Decode(EditModel)
	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	id := u.GetUserIDFromRequest(r)
	Model := models.BeeFarm{Location: EditModel.Location,
		Name: EditModel.Name, BeeFarmTypeID: EditModel.BeeFarmTypeID,
		BeeFarmSizeID: EditModel.BeeFarmSizeID}
	Model.UserID = id

	db := db.GetDB()
	err = db.Create(&Model).Error

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.Respond(w, u.Message(true, "OK"))
	}
}



