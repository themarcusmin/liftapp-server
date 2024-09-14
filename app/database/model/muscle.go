package model

import (
	"gorm.io/gorm"
)

type Muscle struct {
	gorm.Model
	DisplayName string     `json:"displayName,omitempty"`
	Exercise    []Exercise `gorm:"many2many:exercise_muscles" json:"exercises,omitempty"`
}
