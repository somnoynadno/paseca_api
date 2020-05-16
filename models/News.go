package models

type News struct {
	BaseModel
	Title string `json:"text" gorm:"size:512;"`
	Text  string `json:"text" gorm:"size:16384;"`
}
