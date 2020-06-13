package OtherControllers

import (
	"encoding/json"
	"net/http"
	"paseca/db"
	u "paseca/utils"
)

type StatsResponse struct {
	UsersCount       uint
	BeeFarmsCount    uint
	BeeFamiliesCount uint
}

var GetStatsForLanding = func(w http.ResponseWriter, r *http.Request) {
	var usersCount uint
	var beeFarmsCount uint
	var beeFamiliesCount uint

	db := db.GetDB()

	db.Table("users").Count(&usersCount)
	db.Table("bee_farms").Count(&beeFarmsCount)
	db.Table("bee_families").Count(&beeFamiliesCount)

	response := StatsResponse{UsersCount: usersCount, BeeFamiliesCount: beeFamiliesCount, BeeFarmsCount: beeFarmsCount}
	res, err := json.Marshal(response)

	if err != nil {
		u.HandleBadRequest(w, err)
	} else {
		u.RespondJSON(w, res)
	}
}