package models

type Hive struct {
	BaseModel
	Name            string        `json:"name" gorm:"not null;index;"`
	CoordX          *int          `json:"coord_x"`
	CoordY          *int          `json:"coord_y"`
	HiveFormatID    uint          `json:"hive_format_id"`
	HiveFormat      HiveFormat    `json:"hive_format"`
	HiveFrameTypeID uint          `json:"hive_frame_type_id"`
	HiveFrameType   HiveFrameType `json:"hive_frame_type"`
	BeeFarmID       uint          `json:"bee_farm_id"`
	BeeFarm         BeeFarm       `json:"bee_farm"`
	BeeFamilyID     *uint         `json:"bee_family_id"`
	UserID          uint          `json:"user_id"`
}
