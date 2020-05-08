package models

type BeeFarmType struct {
	BaseModel
	Name string `json:"name" gorm:"not null;"`
}
