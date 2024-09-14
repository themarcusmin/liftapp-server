package model

type ExerciseWorkoutJunction struct {
	ID            uint    `gorm:"primaryKey"`
	PrescribedSet *uint8  `json:"prescribedSet,omitempty"`
	PrescribedRep *uint8  `json:"prescribedRep,omitempty"`
	RestTime      *uint16 `json:"prescribedRestTime,omitempty"`
	ExerciseID    uint
	WorkoutID     *uint
}
