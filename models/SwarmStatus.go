package models

type SwarmStatus struct {
	BaseModelWithCustom
	Status string  `json:"status" gorm:"not null;"`
}
