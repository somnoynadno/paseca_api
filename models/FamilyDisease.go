package models

type FamilyDisease struct {
	BaseModelWithUser
	BeeFamilyID  uint       `json:"bee_family_id" gorm:"not null;"`
	BeeFamily    BeeFamily  `json:"bee_family"`
	BeeDiseaseID uint       `json:"bee_disease_id" gorm:"not null;"`
	BeeDisease   BeeDisease `json:"bee_disease"`
}