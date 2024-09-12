package model

// ExerciseMuscleGroup model - intermediate table `exercise_muscles` (many to many relations)
type ExerciseMuscle struct {
	ExerciseID uint `gorm:"primaryKey"`
	MuscleID   uint `gorm:"primaryKey"`
	ISPRIMARY  bool `gorm:"not null" json:"isPrimary"`
}
