package models

type BeeDisease struct {
	BaseModelWithCustom
	Name         string       `json:"name" gorm:"not null;"`
	Descriptions *string      `json:"description" gorm:"size:4096"`
	BeeFamilies  []*BeeFamily `json:"bee_families" gorm:"many2many:family_diseases;"`
}
