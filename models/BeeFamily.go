package models

import "time"

type BeeFamily struct {
	BaseModelWithUser
	Name               string            `json:"name" gorm:"not null;index;"`
	QueenBeeBornDate   *time.Time        `json:"queen_bee_born_date"`
	LastInspectionDate *time.Time        `json:"last_inspection_date"`
	BeeFarmID          uint              `json:"bee_farm_id"`
	BeeFarm            *BeeFarm          `json:"bee_farm"`
	BeeBreedID         uint              `json:"bee_breed_id"`
	BeeBreed           *BeeBreed         `json:"bee_breed"`
	BeeFamilyStatusID  uint              `json:"bee_family_status_id"`
	BeeFamilyStatus    BeeFamilyStatus   `json:"bee_family_status"`
	HiveID             *uint             `json:"hive_id,omitempty"`
	Hive               *Hive             `json:"hive,omitempty"`
	HoneyHarvests      []*HoneyHarvest   `json:"honey_harvests,omitempty"`
	BeeDiseases        []*BeeDisease     `json:"bee_diseases,omitempty" gorm:"many2many:family_diseases;"`
	Parent1ID          *uint             `json:"parent1_id,omitempty"`
	Parent1            *BeeFamily        `json:"parent1,omitempty"`
	Parent2ID          *uint             `json:"parent2_id,omitempty"`
	Parent2            *BeeFamily        `json:"parent2,omitempty"`
	IsControl          bool              `json:"is_control" gorm:"default:false; not null;"`
	ControlHarvests    []*ControlHarvest `json:"control_harvests,omitempty"`
}
