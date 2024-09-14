package model

import (
	"gorm.io/gorm"
)

type ProgramDay struct {
	gorm.Model
	DisplayName string    `json:"displayName"`
	Workout     []Workout `gorm:"foreignKey:ProgramDayID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"workout"`
	ProgramID   uint
}
