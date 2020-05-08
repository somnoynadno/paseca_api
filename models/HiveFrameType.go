package models

type HiveFrameType struct {
	BaseModel
	Name string `json:"name" gorm:"not null;"`
}
