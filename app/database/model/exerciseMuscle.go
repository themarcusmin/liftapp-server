package model

// ExerciseMuscleGroup model - intermediate table `exercise_muscles` (many to many relations)
type ExerciseMuscle struct {
	ExerciseID uint `gorm:"primaryKey"`
	MuscleID   uint `gorm:"primaryKey"`
	IsPRIMARY  bool `gorm:"not null" json:"isPrimary"`
}
