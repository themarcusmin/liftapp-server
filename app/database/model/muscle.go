package model

import (
	"gorm.io/gorm"
)

type Muscle struct {
	gorm.Model
	DisplayName string     `json:"displayName,omitempty"`
	Exercise    []Exercise `gorm:"many2many:exercise_Muscles" json:"exercies,omitempty"`
}
