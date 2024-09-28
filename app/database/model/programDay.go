package model

import (
	"gorm.io/gorm"
)

type ProgramDay struct {
	gorm.Model
	DisplayName     string            `json:"displayName"`
	ProgramExercise []ProgramExercise `gorm:"foreignKey:ProgramDayID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"programExercises"`
	Log             []Log             `gorm:"foreignKey:ProgramDayID;references:ID;" json:"logs"`
	ProgramID       uint
}
