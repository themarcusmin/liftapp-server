package model

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	DisplayName string   `json:"displayName"`
	Muscle      []Muscle `gorm:"many2many:exercise_muscles" json:"muscles,omitempty"`
}
