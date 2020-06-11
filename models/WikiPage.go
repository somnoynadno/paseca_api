package models

type WikiPage struct {
	BaseModel
	Title         string       `json:"title" gorm:"not null; size:255"`
	Description   *string      `json:"description" gorm:"size:4096"`
	Content       string       `json:"content" gorm:"type:text;not null"`
	WikiSectionID uint         `json:"wiki_section_id"`
	WikiSection   WikiSection  `json:"wiki_section,omitempty"`
	Author        *string      `json:"author" gorm:"size:255"`
}
