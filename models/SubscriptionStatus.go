package models

type SubscriptionStatus struct {
	BaseModel
	Status string `json:"status" gorm:"not null;"`
}
