package models

import "time"

type ControlHarvest struct {
	BaseModel
	Date        *time.Time `json:"date"`
	Amount      float64    `json:"amount" gorm:"not null;"`
	BeeFamilyID uint       `json:"bee_family_id"`
	BeeFamily   BeeFamily  `json:"bee_family"`
	UserID      uint       `json:"user_id"`
}
