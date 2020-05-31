package models

import "time"

type Swarm struct {
	BaseModelWithUser
	Date          *time.Time   `json:"date"`
	Order         *int         `json:"order"`
	BeeFamilyID   uint         `json:"bee_family_id"`
	BeeFamily     *BeeFamily   `json:"bee_family,omitempty"`
	SwarmStatusID uint         `json:"swarm_status_id"`
	SwarmStatus   *SwarmStatus `json:"swarm_status,omitempty"`
}
