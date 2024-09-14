package model

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	EventDate  time.Time
	UserID     uint
	WorkoutID  *uint
	LogEntries []LogEntry `gorm:"foreignKey:LogID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"logEntries"`
}
