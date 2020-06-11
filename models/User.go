package models

import "time"

type User struct {
	BaseModel
	Email                string             `json:"email" gorm:"not null;unique_index;"`
	Name                 string             `json:"name" gorm:"not null;"`
	Surname              string             `json:"surname" gorm:"not null;"`
	Password             string             `json:"password" gorm:"not null;"`
	IsAdmin              bool               `json:"is_admin" gorm:"not null; default:false;"`
	LastLogin            *time.Time         `json:"last_login"`
	SubscriptionEnd      *time.Time         `json:"subscription_end"`
	SubscriptionStatusID uint               `json:"subscription_status_id"`
	SubscriptionStatus   SubscriptionStatus `json:"subscription_status,omitempty"`
	SubscriptionTypeID   uint               `json:"subscription_type_id"`
	SubscriptionType     SubscriptionType   `json:"subscription_type,omitempty"`
	BeeFarms             []BeeFarm          `json:"bee_farms,omitempty"`
}
