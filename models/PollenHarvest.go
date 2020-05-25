package models

import "time"

type PollenHarvest struct {
	BaseModelWithUser
	Date        *time.Time `json:"date"`
	Amount      float64    `json:"amount" gorm:"not null;"`
	BeeFarmID   uint       `json:"bee_farm_id"`
	BeeFarm     BeeFarm    `json:"bee_farm,omitempty"`
}
