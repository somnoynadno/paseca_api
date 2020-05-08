package models

import (
	"time"
)

type HoneyHarvest struct {
	BaseModel
	Date        *time.Time `json:"date"`
	Amount      float64    `json:"amount" gorm:"not null;"`
	TotalPrice  float64    `json:"total_price" gorm:"not null;"`
	HoneyTypeID uint       `json:"honey_type_id"`
	HoneyType   HoneyType  `json:"honey_type"`
	BeeFamilyID uint       `json:"bee_family_id"`
	BeeFamily   BeeFamily  `json:"bee_family"`
}
