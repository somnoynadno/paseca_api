package models

type WikiSection struct {
	BaseModel
	Name         string `json:"name" gorm:"not null; size:255;"`
	Description *string `json:"description" gorm:"size:4096"`
	URL          string `json:"url" gorm:"not null; size:255"`
}
