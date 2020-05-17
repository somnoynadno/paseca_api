package models

type HiveFrameType struct {
	BaseModelWithCustom
	Name string `json:"name" gorm:"not null;"`
}
