package models

type SubscriptionType struct {
	BaseModel
	Name        string  `json:"name" gorm:"not null;"`
	Description *string `json:"description"`
	Price       int     `json:"price"`
}
