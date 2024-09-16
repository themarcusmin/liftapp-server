package model

import (
	"time"

	"gorm.io/gorm"
)

type ProgramEntry struct {
	gorm.Model
	SetNumber          uint8     `json:"setNumber"`
	PrescribedReps     *uint8    `json:"prescribedReps"`
	PrescribedDuration *uint16   `json:"prescribedDuration"`
	RestTime           time.Time `json:"restTime"`
	ProgramExerciseID  uint
	LogEntry           []LogEntry `gorm:"foreignKey:ProgramEntryID;references:ID;" json:"logEntries"`
}
