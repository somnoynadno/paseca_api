package models

import "time"

type BaseModel struct {
	ID        uint       `json:"id" gorm:"primary_key" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type BaseModelWithCustom struct {
	ID        uint       `json:"id" gorm:"primary_key" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
	IsCustom  bool       `json:"is_custom" gorm:"not null; default:true;"`
}
