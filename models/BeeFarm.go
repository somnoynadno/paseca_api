package models

type BeeFarm struct {
	BaseModelWithUser
	Name           string           `json:"name" gorm:"not null;index;"`
	Location       *string          `json:"location"`
	User           *User            `json:"user,omitempty"`
	BeeFarmTypeID  uint             `json:"bee_farm_type_id"`
	BeeFarmType    BeeFarmType      `json:"bee_farm_type,omitempty"`
	BeeFarmSizeID  uint             `json:"bee_farm_size_id"`
	BeeFarmSize    BeeFarmSize      `json:"bee_farm_size,omitempty"`
	Reminders      []Reminder       `json:"reminders"`
	BeeFamilies    []BeeFamily      `json:"bee_families"`
	HoneySales     []*HoneySale     `json:"honey_sales,omitempty"`
	Hives          []Hive           `json:"hives"`
	PollenHarvests []*PollenHarvest `json:"pollen_harvests,omitempty"`
}
