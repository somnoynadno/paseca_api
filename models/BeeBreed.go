package models

type BeeBreed struct {
	BaseModelWithCustom
	Name        string `json:"name" gorm:"not null;"`
	Description string `json:"description" gorm:"size:4096;"`
}
