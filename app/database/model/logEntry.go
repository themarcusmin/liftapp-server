package model

import (
	"time"

	"gorm.io/gorm"
)

type LogEntry struct {
	gorm.Model
	SetNumber      uint8     `json:"setNumber"`
	Reps           *uint8    `json:"reps"`
	Weight         *float64  `gorm:"type:decimal(6,2)" json:"weight"`
	Duration       *uint16   `json:"duration"`
	EventAt        time.Time `json:"eventAt"`
	LogExerciseID  uint
	ProgramEntryID *uint
}
