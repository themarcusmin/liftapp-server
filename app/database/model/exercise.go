package model

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	DisplayName     string            `json:"displayName"`
	ProgramExercise []ProgramExercise `gorm:"foreignKey:ExerciseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"programExercises"`
	LogExercise     []LogExercise     `gorm:"foreignKey:ExerciseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"logExercises"`
	FormatID        uint
	Format          Format     `gorm:"foreignKey:FormatID"`
	ParentID        *uint      `json:"parentId"`
	SubExercises    []Exercise `gorm:"foreignKey:ParentID" json:"subExercises,omitempty"`
	Muscles         []Muscle   `gorm:"many2many:exercise_muscles" json:"musclesWorked"`
}
