package model

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	EventAt      time.Time `json:"eventAt"`
	UserID       uint64
	ProgramDayID *uint
	LogExercise  []LogExercise `gorm:"foreignKey:LogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"logExercises"`
}
