package model

import (
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	DisplayName  string     `json:"displayName"`
	Exercises    []Exercise `gorm:"many2many:exercise_workout_junctions;" json:"exercises"`
	Logs         []Log      `gorm:"foreignKey:WorkoutID;references:ID;constraint:OnUpdate:CASCADE" json:"logs"`
	ProgramDayID uint
}
