package model

import (
	"gorm.io/gorm"
)

type LogExercise struct {
	gorm.Model
	LogID             uint
	ExerciseID        uint
	ProgramExerciseID *uint
	LogEntry          []LogEntry `gorm:"foreignKey:LogExerciseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"logEntries"`
}
