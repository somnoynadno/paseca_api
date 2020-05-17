package models

type BeeFarm struct {
	BaseModel
	Name          string       `json:"name" gorm:"not null;index;"`
	Location      *string      `json:"location"`
	UserID        uint         `json:"user_id"`
	User          User         `json:"user"`
	BeeFarmTypeID uint         `json:"bee_farm_type_id"`
	BeeFarmType   BeeFarmType  `json:"bee_farm_type"`
	BeeFarmSizeID uint         `json:"bee_farm_size_id"`
	BeeFarmSize   BeeFarmSize  `json:"bee_farm_size"`
	Reminders     *[]Reminder  `json:"reminders"`
	BeeFamilies   *[]BeeFamily `json:"bee_families"`
	HoneySales    *[]HoneySale `json:"honey_sales"`
	Hives         *[]Hive      `json:"hives"`
}
