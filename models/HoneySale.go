package models

import "time"

type HoneySale struct {
	BaseModel
	Date        *time.Time `json:"date"`
	Amount      float64    `json:"amount" gorm:"not null;"`
	TotalPrice  float64    `json:"total_price" gorm:"not null;"`
	HoneyTypeID uint       `json:"honey_type_id"`
	HoneyType   HoneyType  `json:"honey_type"`
	BeeFarmID   uint       `json:"bee_farm_id"`
	BeeFarm     BeeFarm    `json:"bee_farm"`
	UserID      uint       `json:"user_id"`
}
