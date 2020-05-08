package models

type HiveFormat struct {
	BaseModel
	Name string `json:"name" gorm:"not null;"`
	Size *int   `json:"size"`
}
