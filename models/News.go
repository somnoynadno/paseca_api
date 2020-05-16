package models

type News struct {
	BaseModel
	Title string `json:"title" gorm:"size:512;"`
	Text  string `json:"text" gorm:"size:16384;"`
}
