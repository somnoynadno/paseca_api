package models

type HoneyType struct {
	BaseModel
	Name        string  `json:"name" gorm:"not null;"`
	BasePrice   int     `json:"base_price" gorm:"not null;"`
	Description *string `json:"description" gorm:"size:1024;"`
}
