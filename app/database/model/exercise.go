package model

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	DisplayName string    `json:"displayName"`
	Muscle      []Muscle  `gorm:"many2many:exercise_muscles" json:"muscles,omitempty"`
	Workout     []Workout `gorm:"many2many:exercise_workout_junctions" json:"workouts,omitempty"`
}
