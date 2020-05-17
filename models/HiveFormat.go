package models

type HiveFormat struct {
	BaseModelWithCustom
	Name string `json:"name" gorm:"not null;"`
	Size *int   `json:"size"`
}
