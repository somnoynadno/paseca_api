package models

type BeeFamilyStatus struct {
	BaseModel
	Status string  `json:"status" gorm:"not null;"`
	Color  *string `json:"color"`
}
