package model

import (
	"gorm.io/gorm"
)

type ProgramExercise struct {
	gorm.Model
	ExerciseID   uint
	ProgramDayID uint
	LogExercise  []LogExercise  `gorm:"foreignKey:ProgramExerciseID;references:ID;" json:"logExercises"`
	ProgramEntry []ProgramEntry `gorm:"foreignKey:ProgramExerciseID;references:ID;" json:"programEntries"`
}
