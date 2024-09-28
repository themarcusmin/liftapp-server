package model

import (
	"gorm.io/gorm"
)

type ProgramEntry struct {
	gorm.Model
	ProgramExerciseID         uint
	SetNumber                 uint8      `json:"setNumber"`
	PrescribedReps            *uint8     `json:"reps"`
	PrescribedOneRmPrecentage *uint8     `json:"oneRmPercentage"`
	PrescribedDuration        *uint16    `json:"duration"`
	LogEntry                  []LogEntry `gorm:"foreignKey:ProgramEntryID;references:ID;" json:"logEntries"`
}
