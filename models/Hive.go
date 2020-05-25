package models

type Hive struct {
	BaseModelWithUser
	Name            string        `json:"name" gorm:"not null;index;"`
	CoordX          *int          `json:"coord_x"`
	CoordY          *int          `json:"coord_y"`
	HiveFormatID    uint          `json:"hive_format_id"`
	HiveFormat      HiveFormat    `json:"hive_format,omitempty"`
	HiveFrameTypeID uint          `json:"hive_frame_type_id"`
	HiveFrameType   HiveFrameType `json:"hive_frame_type,omitempty"`
	BeeFarmID       uint          `json:"bee_farm_id"`
	BeeFarm         *BeeFarm      `json:"bee_farm,omitempty" gorm:"PRELOAD:false"`
	BeeFamilyID     *uint         `json:"bee_family_id"`
}
