package model

import (
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	DisplayName  string     `json:"displayName"`
	Exercise     []Exercise `gorm:"many2many:exercise_workout_junctions;" json:"exercises"`
	ProgramDayID uint
}
