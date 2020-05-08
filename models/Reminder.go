package models

import "time"

type Reminder struct {
	BaseModel
	Title     string    `json:"title" gorm:"not null;"`
	Text      string    `json:"text" gorm:"not null;"`
	Date      time.Time `json:"date"`
	IsChecked bool      `json:"is_checked" gorm:"not null;default:false;"`
	BeeFarmID uint      `json:"bee_farm_id"`
	BeeFarm   BeeFarm   `json:"bee_farm"`
}
