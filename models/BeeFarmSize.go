package models

type BeeFarmSize struct {
	BaseModel
	Name        string  `json:"name" gorm:"not null"`
	Description *string `json:"description"`
	MaxX        int     `json:"max_x" gorm:"not null"`
	MaxY        int     `json:"max_y" gorm:"not null"`
}
