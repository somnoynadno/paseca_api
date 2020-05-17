package models

type BeeFarmType struct {
	BaseModelWithCustom
	Name string `json:"name" gorm:"not null;"`
}
