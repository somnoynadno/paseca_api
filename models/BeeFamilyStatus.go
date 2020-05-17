package models

type BeeFamilyStatus struct {
	BaseModelWithCustom
	Status string  `json:"status" gorm:"not null;"`
	Color  *string `json:"color"`
}
