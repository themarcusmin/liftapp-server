package model

import (
	"gorm.io/gorm"
)

type ProgramExercise struct {
	gorm.Model
	ExerciseID   uint `json:"exerciseID"`
	ProgramDayID uint
	RestTime     uint8          `json:"restTime"`
	LogExercise  []LogExercise  `gorm:"foreignKey:ProgramExerciseID;references:ID;" json:"logExercises"`
	ProgramEntry []ProgramEntry `gorm:"foreignKey:ProgramExerciseID;references:ID;" json:"programEntries"`
}
