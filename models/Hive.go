package models

type Hive struct {
	BaseModel
	Name            string        `json:"name" gorm:"not null;index;"`
	CoordX          *int          `json:"coord_x"`
	CoordY          *int          `json:"coord_y"`
	BeeFamilyID     *uint         `json:"bee_family_id"`
	BeeFamily       *BeeFamily    `json:"bee_family"`
	HiveFormatID    uint          `json:"hive_format_id"`
	HiveFormat      HiveFormat    `json:"hive_format"`
	HiveFrameTypeID uint          `json:"hive_frame_type_id"`
	HiveFrameType   HiveFrameType `json:"hive_frame_type"`
}
