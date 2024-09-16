package model

import (
	"gorm.io/gorm"
)

type ProgramDay struct {
	gorm.Model
	DisplayName string `json:"displayName"`
	Log         []Log  `gorm:"foreignKey:ProgramDayID;references:ID;" json:"logs"`
	ProgramID   uint
}
