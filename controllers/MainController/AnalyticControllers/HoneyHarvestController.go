package AnalyticControllers

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	u "paseca/utils"
)

type HoneyHarvestsGroupByAmount struct {
	BeeFamilyID   uint    `json:"bee_family_id"`
	BeeFamilyName string  `json:"bee_family_name"`
	Amount        float64 `json:"amount"`
}

type HoneyHarvestsResponse struct {
	BeeFarmID     uint   `json:"bee_farm_id"`
	BeeFarmName   string `json:"bee_farm_name"`
	HoneyHarvests []HoneyHarvestsGroupByAmount `json:"honey_harvests"`
}

var GetHoneyHarvestsGroupByAmount = func(w http.ResponseWriter, r *http.Request) {
	var entities []HoneyHarvestsResponse
	id := r.Context().Value("context").(u.Values).Get("user_id")

	db := db.GetDB()
	err := db.Table("bee_farms").
		Select("id as bee_farm_id, name as bee_farm_name").
		Where("user_id = ? and deleted_at is null", id).Scan(&entities).Error

	if err != nil {
		u.HandleBadRequest(w, err)
		return
	}

	for i := 0; i < len(entities); i++ {
		var harvests []HoneyHarvestsGroupByAmount

		err := db.Table("honey_harvests").
			Joins("join bee_families on bee_families.id = honey_harvests.bee_family_id").
			Select("bee_families.id as bee_family_id, bee_families.name as bee_family_name, sum(amount) as amount").
			Group("bee_families.id").
			Where("bee_families.bee_farm_id = ? and honey_harvests.deleted_at is null", entities[i].BeeFarmID).
			Scan(&harvests).Error

		if err != nil {
			u.HandleBadRequest(w, err)
			return
		}

		entities[i].HoneyHarvests = harvests
	}

	res, err := json.Marshal(entities)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}
